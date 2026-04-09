package service

import (
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Translate(data string) (string, error) {
	cleanInput := strings.TrimSpace(data)
	isText := false

	for _, char := range cleanInput {
		if unicode.IsLetter(char) {
			isText = true
			break
		}
	}

	if isText {

		return morse.DefaultConverter.ToMorse(cleanInput), nil
	} else {

		return morse.DefaultConverter.ToText(cleanInput), nil
	}

}
