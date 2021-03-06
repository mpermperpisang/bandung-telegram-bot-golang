package helper

import (
	"regexp"
)

var staging string

func CheckEmptyStaging(text string) string {
	pattern := regexp.MustCompile(RegexCompileCheckEmptyStaging())
	staging = pattern.FindString(text)

	return staging
}

func CheckEmptySquadStaging(text string) string {
	pattern := regexp.MustCompile(RegexCompileCheckEmptySquadStaging())
	staging = pattern.FindString(text)

	return staging
}
