package message

func OnCall(backend string) string {
	return "Kak " + backend + " oncall hari ini\nSemangat yaa 😘"
}

func AddOnCall(year string) string {
	return "Berhasil mengupdate data BE oncall untuk tahun " + year
}

func HolidayOnCall() string {
	return "Nabilah lagi liburan dulu yaa, Kak ✈️🏝⛱\n" +
		"Sampai ketemu di hari kerja 👋🏻👋🏻👋🏻"
}
