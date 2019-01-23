package message

func EmptyStaging(username string) string {
	return "Nama stagingnya ndak ada, Kak @" + username + "\n" +
		"Contoh : /status_staging 21 51 103 atau /status_staging DANA"
}

func StagingNotFound(staging string) string {
	return "Staging<b>" + staging + "</b> tidak ditemukan, Kak"
}
