package message

func EmptyStaging(username string) string {
	return "Nama stagingnya ndak ada, Kak @" + username + "\n" +
		"Contoh : /status_staging staging_number1 staging_number2 staging_number3\n" +
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
		"Contoh : " + command + " squad_name staging_number1 staging_number2 staging_number3"
}

func EmptyOncallUsername(username string) string {
	return "Commandnya belum benar, Kak @" + username + "\n" +
		"Contoh : /add_oncall @username1 @username2 @username3"
}

func HappyNewYear() string {
	return "Selamat Tahun Baru, Kakak-kakak\n" +
		"Mau ngingetin nih, jadwal BE oncallnya tolong diperbarui yaa\n" +
		"Happy holiday 🥳🎉"
}
