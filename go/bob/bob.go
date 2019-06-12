// bob implements lackadaisical teenager
package bob

import (
	"strings"
)

// Hey returns Bob's answers
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)
	isQuestion := strings.HasSuffix(remark, "?")
	isUpper := strings.ToUpper(remark) == remark
	isNoLetters := strings.ToUpper(remark) == remark && strings.ToLower(remark) == remark

	if isUpper && isQuestion && !isNoLetters {
		return "Calm down, I know what I'm doing!"
	} else if isUpper && !isNoLetters {
		return "Whoa, chill out!"
	} else if isQuestion {
		return "Sure."
	} else if remark == "" {
		return "Fine. Be that way!"
	}

	return "Whatever."
}
