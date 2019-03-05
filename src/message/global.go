package message

func EmptyStatusStaging(username, command string) string {
	return "Nama stagingnya ndak ada, Kak @" + username + "\n" +
		"Contoh : " + command + " staging_number1 staging_number2 staging_number3 dan seterusnya\n" +
		"atau /status_staging squad_name"
}

func EmptyBookDoneStaging(username, command string) string {
	return "Nama stagingnya ndak ada, Kak @" + username + "\n" +
		"Contoh : " + command + " staging_number"
}

func StagingNotFound(staging string) string {
	return "staging<b>" + staging + ".vm</b> tidak ditemukan, Kak"
}

func EmptySquadStaging(username, command string) string {
	return "Commandnya belum benar, Kak @" + username + "\n" +
		"Contoh : " + command + " squad_name staging_number1 staging_number2 staging_number3 dan seterusnya"
}

func EmptyDayUsername(username, command string) string {
	return "Commandnya belum benar, Kak @" + username + "\n" +
		"Contoh : " + command + "day1 @username @day2 @username @day3 @username dan seterusnya"
}

func EmptyUsername(username, command string) string {
	return "Commandnya belum benar, Kak @" + username + "\n" +
		"Contoh : " + command + " @username1 @username2 @username3 dan seterusnya"
}
