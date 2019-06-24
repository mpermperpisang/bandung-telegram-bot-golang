package message

import (
	"io/ioutil"
	"strings"
)

func AddSnack(snack, username string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data) + "\n\nCatat juga di <a href='https://bit.ly/2FBKhA4'>CONFLUENCE</a> yaa\n" +
		"Dan baca juga aturan per-snack-an di link tersebut 🙏🏻"

	if strings.Contains(content, "bawa snack") {
		header = "Cihuy Kak @" + username + " nambahin jadwal\n"
	} else {
		header = "Kok didaftarin lagi sih, Kak @" + username + "? 😵\n"
	}

	return header + content
}

func MoveSnack(snack, username string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if strings.Contains(content, "pindah sementara") {
		header = "Cihuy Kak @" + username + " mindahin jadwal\n"
	} else {
		header = "Mindahin siapa sih, Kak @" + username + "? 😵\n"
	}

	return header + content
}

func PermanentSnack(snack, username string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if strings.Contains(content, "pindah selamanya") {
		header = "Cihuy Kak @" + username + " mindahin jadwal\n"
	} else {
		header = "Mindahin siapa sih, Kak @" + username + "? 😵\n"
	}

	return header + content
}

func DeleteSnack(snack, username string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if strings.Contains(content, "ByBy") {
		header = "Colek Kak @" + username + ", tolong hapus data Kak <code>" + snack + "</code> di tiny.cc/bukasnack yaa 😙\n"
	} else {
		header = "Ngehapus siapa sih, Kak @" + username + "? 😵\n"
	}

	return header + content
}

func DoneSnack(username, first_name string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if strings.Contains(content, "menggendutkan") {
		header = "Yeay dapat cemilan dari Kak @" + username + ""
	} else {
		header = "Nge-done-in siapa sih, Kak " + first_name + "? 😵\n"
	}

	return header + content
}

func CancelSnack(snack, username string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if strings.Contains(content, "ora jadi") {
		header = "Snacknya dibatalin dulu yaa\n"
	} else {
		header = "Ngebatalin siapa sih, Kak @" + username + "? 😵\n"
	}

	return header + content
}

func HolidaySnack(snack, username string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if strings.Contains(content, "izin libur") {
		header = "Snacknya diliburin dulu yaa\n"
	} else if strings.Contains(content, "Selamat") {
		header = ""
	} else {
		header = "Ngeliburin siapa sih, Kak @" + username + "? 😵\n" +
			"Hanya bisa ngeliburan orang yang jadwalnya hari ini ajah soale\n"
	}

	return header + content
}

func AddAdminSnack(snack, username string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if strings.Contains(content, "resmi") {
		header = "Cihuy Kak @" + username + " nambahin admin snack 🎉👏🏻🥳\n"
	} else {
		header = "Kok adminnya didaftarin lagi sih, Kak @" + username + "? 😵\n"
	}

	return header + content
}

func DeleteAdminSnack(snack, username string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if strings.Contains(content, "dihapus") {
		header = "Yaah berkurang deh adminnya, Kak @" + username + "\n"
	} else {
		header = "Ngehapus admin siapa sih, Kak @" + username + "? 😵\n"
	}

	return header + content
}

func ListAdminSnack(username string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if content != "" {
		header = "Berikut adalah daftar list admin snack, Kak @" + username + "\n"
	} else {
		header = "Maaf Kak @" + username + " belum ada list admin snack\n"
	}

	return header + content
}

func BaikSnack(username string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if content != "" {
		header = "Berikut adalah Kakak yang baik hati sering berbagi snack\n\n"
	} else {
		header = "Maaf Kak @" + username + " belum ada list Kakak yang baik hati\n"
	}

	return header + content
}

func AskSnack(username, text string) string {
	var header string

	data, _ := ioutil.ReadFile("temp.go")
	content := string(data)

	if content != "" {
		header = "Hi Kak.\nKak @" + username + " nanya snack<b>" + text + "</b> siapa yang beli dan dimana yaa? Enak cenah 😍😍😍\n" +
			"Tolong kabari langsung ke yg nanya yaa 😘"
	} else {
		header = "Belum ada yang bawa snack"
	}

	return header
}
