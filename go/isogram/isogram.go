package isogram

import (
	"regexp"
	"strings"
)

// Determine whether a sentence has any duplicate letters
func IsIsogram(sentence string) bool {
	regex := regexp.MustCompile(`[^a-z]`)
	letters := regex.ReplaceAllString(strings.ToLower(sentence), "")

	for i := 0; i < len(letters) - 1; i += 1 {
		if (1 != strings.Count(letters, letters[i : i+1])) {
			return false
		}
	}

	return true
}
