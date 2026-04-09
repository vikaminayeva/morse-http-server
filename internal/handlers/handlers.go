package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, req *http.Request) {
	file, header, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error ride", http.StatusInternalServerError)
		return
	}
	content := string(data)
	result, err := service.Translate(content)
	if err != nil {
		http.Error(w, "Ошибка перевода", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(header.Filename)
	fileName := time.Now().UTC().String() + ext

	newFile, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Error create", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	_, err = newFile.WriteString(result)
	if err != nil {
		http.Error(w, "Errors write", http.StatusInternalServerError)
		return
	}
	w.Write([]byte(result))
}
