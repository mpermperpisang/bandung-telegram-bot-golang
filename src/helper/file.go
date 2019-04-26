package helper

import "os"

func CreateFile() (file *os.File) {
	file, err := os.Create("temp.go")
	ErrorMessage(err)

	return file
}

func OpenFile() (file *os.File) {
	file, err := os.Create("temp.go")
	ErrorMessage(err)

	return file
}
