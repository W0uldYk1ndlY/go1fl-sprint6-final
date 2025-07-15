package service

import "github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"

func DetermineMorseOrNot(s string) string {

	for _, char := range s {
		if char != '.' && char != '-' && char != ' ' {
			result := morse.ToMorse(s)
			return result
		}
	}
	result := morse.ToText(s)
	return result
}
