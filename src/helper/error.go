package helper

import (
	"fmt"
)

type error interface {
	Error() string
}

func ErrorMessage(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
