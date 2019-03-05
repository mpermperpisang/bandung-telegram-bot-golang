package message

func DefaultHelp() string {
	return "Maaf, Kak\nKonten bantuan belum tersedia"
}

func OnlineMessage(name string) string {
	return name + " : <b>ONLINE</b>" +
		"\n\nJangan lupa semangat yaa, Kak. 🥰"
}

func UserSpammer(name string) string {
	return "Kak @" + name + " udah terlalu banyak kirim command yang sama\n" +
		"Gantian sama yang lain yaa biar ga dianggap <b>SPAMMER</b>"
}
