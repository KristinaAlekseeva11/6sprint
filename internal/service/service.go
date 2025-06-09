package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ConvertAuto(input string) (string, error) {
	input = strings.TrimSpace(input)

	if input == "" {
		return "", errors.New("input is empty")
	}

	if strings.ContainsAny(input, ".-") && !strings.ContainsAny(input, "АаБбВв") {
		return morse.ToText(input), nil
	}

	return morse.ToMorse(input), nil
}
