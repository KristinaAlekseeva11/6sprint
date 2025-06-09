package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func isMorse(input string) bool {
	allowed := ".- \n"
	for _, ch := range input {
		if !strings.ContainsRune(allowed, ch) {
			return false
		}
	}
	return true
}

func ConvertAuto(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "", errors.New("input is empty")
	}

	if isMorse(input) {
		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}
