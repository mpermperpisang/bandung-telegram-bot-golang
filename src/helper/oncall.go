package helper

import (
	"regexp"
)

var username string

func CheckEmptyUsername(text string) string {
	pattern := regexp.MustCompile(RegexCompileUsername())
	username = pattern.FindString(text)

	return username
}
