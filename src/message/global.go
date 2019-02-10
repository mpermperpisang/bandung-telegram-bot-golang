package message

func EmptyStaging(username string) string {
	return "Nama stagingnya ndak ada, Kak @" + username + "\n" +
		"Contoh : /status_staging staging_number1 staging_number2 staging_number3 atau /status_staging squad_name"
}

func EmptyBookDoneStaging(username string) string {
	return "Nama stagingnya ndak ada, Kak @" + username + "\n" +
		"Contoh : /booking staging_number atau /done staging_number"
}

func StagingNotFound(staging string) string {
	return "Staging<b>" + staging + "</b> tidak ditemukan, Kak"
}

func EmptySquadStaging(username string) string {
	return "Commandnya belum benar, Kak @" + username + "\n" +
		"Contoh : /add_staging squad_name staging_number1 staging_number2 staging_number3"
}
