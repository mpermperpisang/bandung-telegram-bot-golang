package message

func DefaultHelp() string {
	return "Maaf, Kak\nKonten bantuan belum tersedia"
}

func DefaultCommand() string {
	return "Ngetik apa sih, Kak?"
}

func OnlineMessage(name string) string {
	return name + " : <b>ONLINE</b>" +
		"\n\nJangan lupa semangat yaa, Kak. 🥰"
}
