package helper

import (
	"regexp"
)

var username string

func CheckEmptyBackEndOnCall(text string) string {
	pattern := regexp.MustCompile(RegexCompileBackEndOnCall())
	staging = pattern.FindString(text)

	return username
}
