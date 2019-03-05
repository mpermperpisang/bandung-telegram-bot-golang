package helper

import (
	"regexp"
)

var snack string

func CheckEmptyDayUsername(text string) string {
	pattern := regexp.MustCompile(RegexCompileCheckEmptyDayUsername())
	snack = pattern.FindString(text)

	return snack
}
