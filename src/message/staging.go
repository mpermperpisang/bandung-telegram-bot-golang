package message

import (
	"io/ioutil"
)

func AddStaging(staging string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if content != "" {
		header = "Status :"
	} else {
		header = StagingNotFound(staging)
	}

	return header + content
}

func UpdateStaging(staging string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if content != "" {
		header = "Staging berhasil diubah ke squad "
	} else {
		header = StagingNotFound(staging)
	}

	return header + content
}

func BookStaging(first_name, staging string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if content != "" {
		header = "Kak @" + first_name + " memesan "
	} else {
		header = StagingNotFound(staging)
	}

	return header + content
}

func DoneStaging(first_name, staging string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if content != "" {
		header = "Kak @" + first_name + " sudah selesai pakai "
	} else {
		header = StagingNotFound(staging)
	}

	return header + content
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
