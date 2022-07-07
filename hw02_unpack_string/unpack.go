package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var byteArray = []rune(input)
	var length = len(byteArray)
	var result strings.Builder
	var lastRune rune

	if input == "" {
		return result.String(), nil
	}

	for i := 0; i < length; i++ {
		currentRune := byteArray[i]

		if lastRune == 0 {
			if unicode.IsDigit(currentRune) {
				return "", ErrInvalidString
			}

			lastRune = currentRune
			continue
		}

		if unicode.IsDigit(currentRune) && unicode.IsDigit(lastRune) {
			return "", ErrInvalidString
		}

		if unicode.IsDigit(lastRune) && unicode.IsLetter(currentRune) {
			lastRune = currentRune
			continue
		}

		if !unicode.IsDigit(currentRune) {
			result.WriteRune(lastRune)
			lastRune = currentRune
			continue
		}

		currentSymbolNumber, _ := strconv.Atoi(string(currentRune))

		for j := 0; j < currentSymbolNumber; j++ {
			result.WriteRune(lastRune)
		}

		lastRune = currentRune
	}

	lastRune = byteArray[length-1]

	if unicode.IsLetter(lastRune) {
		result.WriteRune(lastRune)
	}

	return result.String(), nil
}
