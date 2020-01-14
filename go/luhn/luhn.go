package luhn

import (
	"regexp"
	"strconv"
)

// Check whether a given string is Valid using the Luhn algorithm
func Valid(number string) bool {
	if !containsOnlyNumbersAndSpaces(number) {
		return false
	}

	number = sanitizeNumber(number)

    if 1 >= len(number) {
    	return false
    }

    return 0 == calculateCheckSum(number)
}

// Check whether the string contains anything except numbers and spaces
func containsOnlyNumbersAndSpaces(number string) bool {
	regex := regexp.MustCompile(`[^0-9 ]`)

	return false == regex.Match([]byte(number))
}

// Remove everything except numbers from string
func sanitizeNumber(number string) string {
	regex := regexp.MustCompile(`[^0-9]`)

	return regex.ReplaceAllString(number, "")
}

// Calculate checksum using Luhn algorithm
func calculateCheckSum(number string) int {
	sum := 0

	for i := 0; i < len(number); i += 1 {
		value, _ := strconv.Atoi(number[len(number) - i - 1 : len(number) - i])

		if 1 == i % 2 {
			value = value * 2

			if value >= 10 {
				value -= 9
			}
		}

		sum += value
	}

	return sum % 10
}
