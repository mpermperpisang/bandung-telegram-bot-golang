package helper

func RegexCompileCheckEmptyStaging() string {
	return `\s[a-zA-Z0-9_ ]+`
}

func RegexCompileStatusStaging() string {
	return `\s[a-zA-Z0-9_]+`
}
