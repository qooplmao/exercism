package bob

import (
	"regexp"
	"strings"
)

func Hey(remark string) string {
	remark = strings.Trim(remark, " \t\n\r")

	if isSilence(remark) {
		return "Fine. Be that way!"
	}

	if isQuestion(remark) {
		if isYell(remark) {
			return "Calm down, I know what I'm doing!"
		}

		return "Sure."
	}

	if isYell(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func isSilence(remark string) bool {
	return "" == remark
}

func isQuestion(remark string) bool {
	return "?" == remark[len(remark)-1:]
}

func isYell(remark string) bool {
	return false == hasMatches("[a-z]", remark) && hasMatches("[A-Z]", remark)
}

func hasMatches(regex string, remark string) bool {
	match, _ := regexp.MatchString(regex, remark)

	return match
}
