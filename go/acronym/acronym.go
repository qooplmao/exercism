package acronym

import (
	"strings"
	"regexp"
)

// Abbreviate a sentence to an acronym, eg. "Central Intelligence Agency" will produce "CIA"
func Abbreviate(s string) string {
	acronym := ""
	alphaNumericAndSpaceRegex := regexp.MustCompile(`[^0-9A-Za-z ]`)

	for _, value := range strings.FieldsFunc(s, split) {
		value = strings.Trim(alphaNumericAndSpaceRegex.ReplaceAllString(value, " "), " ")

		if value != "" {
			acronym = acronym + strings.ToUpper(value[0:1])
		}
	}

	return acronym
}

func split(r rune) bool {
    return r == ' ' || r == '-'
}
