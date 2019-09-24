// Package bob implements lackadaisical teenager.
package bob

import (
	"strings"
)

// Hey returns Bob's answers.
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isQuestion := strings.HasSuffix(remark, "?")
	isUpper := strings.ToUpper(remark) == remark
	isNoLetters := isUpper && strings.ToLower(remark) == remark

	switch {
	case isUpper && isQuestion && !isNoLetters:
		return "Calm down, I know what I'm doing!"
	case isUpper && !isNoLetters:
		return "Whoa, chill out!"
	case isQuestion:
		return "Sure."
	case remark == "":
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
