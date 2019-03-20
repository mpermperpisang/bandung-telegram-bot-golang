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

func RegexCompileNickName() string {
	return `.*(Nama panggilan :).*[a-zA-Z]`
}

func RegexCompileJobTitle() string {
	return `.*(Job title :).*[a-zA-Z]`
}

func RegexCompileSquad() string {
	return `.*(Squad :).*[a-zA-Z]`
}

func RegexCompileLastJob() string {
	return `.*(Pekerjaan atau Pendidikan terakhir :).*[a-zA-Z]`
}

func RegexCompileStatus() string {
	return `(.*(Status :).*[a-zA-Z])`
}

func RegexCompileHobby() string {
	return `(.*(Hobi :).*[a-zA-Z])`
}

func RegexCompileMotto() string {
	return `(.*(Motto :).*[a-zA-Z])`
}
