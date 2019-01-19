package helper

import "os"

func CreateFile() (file *os.File) {
	file, err := os.Create("temp.go")

	if err != nil {
		return
	}

	return file
}
