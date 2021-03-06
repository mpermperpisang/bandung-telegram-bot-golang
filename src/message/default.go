package message

import "strings"

func DefaultHelp() string {
	return "Maaf, Kak\nKonten bantuan belum tersedia"
}

func OnlineMessage(name string) string {
	return name + " : <b>ONLINE</b>" +
		"\n\nJangan lupa semangat yaa, Kak. 🥰"
}

func UserSpammer(username string) string {
	return "Kak @" + username + " udah terlalu banyak kirim command yang sama\n" +
		"Gantian sama yang lain yaa biar ga dianggap <b>SPAMMER</b>"
}

func UserAdmin(username, command string) string {
	return "Maaf Kak @" + username + ", command <code>" + command + "</code> cuma bisa sama admin.\n" +
		"Klik /list_admin yaa"
}

func OncallAdmin(username, command string, admin []string) string {
	adminList := strings.Join(admin, " ")
	return "Maaf Kak @" + username + ", command <code>" + command + "</code> cuma bisa sama admin.\n" +
		"Admin oncall adalah " + adminList
}

func AddedGroup(group string) string {
	return "Halo Kakak-kakak di grup <b>" + group + "</b>\nAsyik grupku bertambah banyak 🥳🥳🥳\n\n" +
		"Kalau butuh bantuan jangan sungkan ketik <b>/help</b> by private message yaa ☺️"
}
