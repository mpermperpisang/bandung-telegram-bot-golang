package message

func OnCall(backend string) string {
	return "Kak " + backend + " oncall hari ini\nSemangat yaa 😘"
}

func AddOnCall(year string) string {
	return "Berhasil mengupdate data BE oncall untuk tahun " + year
}

func HolidayOnCall() string {
	return "Maaf lagi liburan dulu yaa, Kak ✈️🏝⛱\n" +
		"Sampai ketemu di lain hari 👋🏻👋🏻👋🏻"
}

func EmptyOnCall(year string) string {
	return "Data BE oncall untuk tahun " + year + " belum tersedia di tiny.cc/danaoncall\n" +
		"Silahkan hubungi PM atau APM"
}

func EmptyTabSheet(year string) string {
	return "Tab sheet untuk tahun " + year + " tidak ditemukan\n" +
		"Tolong baca aturan penambahan jadwal oncall di https://bit.ly/2GzKZlt"
}
