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

func CheckEmptyAskSnack(text string) string {
	pattern := regexp.MustCompile(RegexCompileAskSnack())
	snack = pattern.FindString(text)

	return snack
}
