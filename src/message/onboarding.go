package message

import (
	"os"
)

func UserJoin(group string, username string) string {
	return "Selamat datang di squad Bandung, Kak @" + username + "\n" +
		"Salam kenal, namaku Shani Indira Natio 🤗\n\n" +
		"<b>CEK JAPRIANKU YAA</b> 😁 (Kakak harus japri @" + os.Getenv("BOT_SNACK") + " duluan dan klik Start yaa)\n\n" +
		"Ada info penting terkait Bukalapak Bandung di message yang ku kirim ke Kakak soalnya\n" +
		"Hatur tengkyu, Kak 🙏🏻"
}

func UserLeft(username string) string {
	return "Sayonara Kak @" + username + "\n" +
		"Jangan lupa sama squad Bandung yaa 👋🏻👋🏻👋🏻"
}

func OnboardingUser(username string) string {
	return "Kak #{name}, yuk perkenalan dulu. Tolong tuliskan biodata dengan format seperti ini :" +
		"🐾 Nama panggilan : " +
		"🐾 Job title : " +
		"🐾 Squad : " +
		"🐾 Pekerjaan atau Pendidikan terakhir : " +
		"🐾 Status : " +
		"🐾 Hobi : " +
		"🐾 Motto : "
}
