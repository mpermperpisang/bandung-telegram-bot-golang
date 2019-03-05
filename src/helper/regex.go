package helper

func RegexCompileCheckEmptyStaging() string {
	return `\s[a-zA-Z0-9_ ]+`
}

func RegexCompileCheckEmptySquadStaging() string {
	return `\s[a-zA-Z]+[0-9_ ]+`
}

func RegexCompileStatusStaging() string {
	return `\s[a-zA-Z0-9_]+`
}

func RegexCompileStagingNumber() string {
	return `\d+`
}

func RegexCompileStagingSquad() string {
	return `[a-zA-Z_]+`
}

func RegexCompileUsername() string {
	return `\B@\S+`
}

func RegexCompileBaseCommand() string {
	return `\/[a-z]\w+`
}

func RegexCompileCheckEmptyDayUsername() string {
	return `\s[a-zA-Z@ ]+`
}

func RegexCompileSnackDay() string {
	return `\s[a-zA-Z]{3}`
}

func RegexCompileSnackSchedule() string {
	return `\s[a-zA-Z0-9]{0}[a-zA-Z0][^\s]+\s@[a-zA-Z0-9_^]+`
}
