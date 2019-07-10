package message

import "strings"

func ExampleCommandStaging(command string) string {
	return "\nContoh : <code>" + command + " squad_name staging_number1 staging_number2 staging_number3</code> dan seterusnya"
}

func EmptyStatusStaging(username, command string) string {
	return "Nama stagingnya ndak ada, Kak @" + username + "\n" +
		"Contoh : <code>" + command + " staging_number1 staging_number2 staging_number3</code> dan seterusnya\n" +
		"atau <code>/status_staging squad_name</code>"
}

func EmptyBookDoneStaging(username, command string) string {
	return "Nama stagingnya ndak ada, Kak @" + username + "\n" +
		"Contoh : <code>" + command + " staging_number</code>"
}

func StagingNotFound(staging string) string {
	return "Staging <b>" + strings.ToUpper(staging) + "</b> tidak ditemukan, Kak. Tambahkan dulu ke daftar yaa"
}

func EmptySquadStaging(username, command string) string {
	return "Commandnya belum benar, Kak @" + username +
		ExampleCommandStaging(command)
}

func EmptyDayUsername(username, command string) string {
	return "Commandnya belum benar, Kak @" + username + "\n" +
		"Contoh : <code>" + command + " day1 @username1 @day2 @username2 @day3 @username3</code> dan seterusnya"
}

func EmptyUsername(username, command string) string {
	return "Commandnya belum benar, Kak @" + username + "\n" +
		"Contoh : <code>" + command + " @username1 @username2 @username3</code> dan seterusnya"
}

func EmptyAskSnack(username, command string) string {
	return "Mau nanya snack yang mana, Kak @" + username + "?\n" +
		"Contoh : <code>" + command + " snack_name1 snack_name2 snack_name3</code> dan seterusnya"
}

func AskingSnack(username string) string {
	return "Sedang ditanyakan ke Kakak-kakak yang sudah bawa snack hari ini yaa, Kak @" + username + ". Tunggu response dari mereka"
}
