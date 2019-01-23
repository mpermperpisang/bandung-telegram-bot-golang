package message

func OnCall(backend string) string {
	return "Kak " + backend + " oncall hari ini\nSemangat yaa 😘"
}

func AddOnCall(year string) string {
	return "Berhasil mengupdate data BE oncall untuk tahun " + year
}
