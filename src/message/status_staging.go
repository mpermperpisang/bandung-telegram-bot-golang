package message

import (
	"io/ioutil"
)

func EmptyStaging(username string) string {
	return "Nama stagingnya ndak ada, Kak @" + username + "\n" +
		"Contoh : /status_staging 21 51 103 atau /status_staging DANA"
}

func StagingStatus(staging string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")

	content := string(data)

	if content != "" {
		header = "Status staging\n\n"
	} else {
		header = StagingNotFound(staging)
	}

	return header + content
}

func StagingNotFound(staging string) string {
	return "Staging<b>" + staging + "</b> tidak ditemukan, Kak"
}
