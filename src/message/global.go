package message

import "strings"

func EmptyStatusStaging(username, command string) string {
	return "Nama stagingnya ndak ada, Kak @" + username + "\n" +
		"Contoh : <code>" + command + " staging_number1 staging_number2 staging_number3</code> dan seterusnya\n" +
		"atau /status_staging squad_name"
}

func EmptyBookDoneStaging(username, command string) string {
	return "Nama stagingnya ndak ada, Kak @" + username + "\n" +
		"Contoh : <code>" + command + " staging_number</code>"
}

func StagingNotFound(staging string) string {
	return "Staging <b>" + strings.ToUpper(staging) + "</b> tidak ditemukan, Kak"
}

func EmptySquadStaging(username, command string) string {
	return "Commandnya belum benar, Kak @" + username + "\n" +
		"Contoh : <code>" + command + " squad_name staging_number1 staging_number2 staging_number3</code> dan seterusnya"
}

func EmptyDayUsername(username, command string) string {
	return "Commandnya belum benar, Kak @" + username + "\n" +
		"Contoh : <code>" + command + " day1 @username1 @day2 @username2 @day3 @username3</code> dan seterusnya"
}

func EmptyUsername(username, command string) string {
	return "Commandnya belum benar, Kak @" + username + "\n" +
		"Contoh : <code>" + command + " @username1 @username2 @username3</code> dan seterusnya"
}
