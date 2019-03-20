package message

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

func AddedGroup(group string) string {
	return "Halo Kakak-kakak di grup <b>" + group + "</b>\nAsyik grupku bertambah banyak 🥳🥳🥳\n\n" +
		"Kalau butuh bantuan jangan sungkan ketik /help by private message yaa ☺️"
}
