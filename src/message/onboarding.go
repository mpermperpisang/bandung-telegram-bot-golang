package message

import (
	"os"
)

func UserJoin(group string, username string) string {
	return "Selamat datang di squad Bandung, Kak @" + username + "\n" +
		"Salam kenal, namaku Shani Indira Natio 🤗\n\n" +
<<<<<<< HEAD
		"<b>CEK JAPRIANKU YAA</b> 😁 (Kakak harus japri @" + os.Getenv("BOT_SNACK") + " duluan dan klik START yaa)\n\n" +
=======
		"<b>CEK JAPRIANKU YAA</b> 😁 (Kakak harus japri @" + os.Getenv("BOT_SNACK") + " duluan dan klik Start yaa)\n\n" +
>>>>>>> 0bf34c0d62a747c009237eed2cf0ede2388788f7
		"Ada info penting terkait Bukalapak Bandung di message yang ku kirim ke Kakak soalnya\n" +
		"Hatur tengkyu, Kak 🙏🏻"
}

<<<<<<< HEAD
func UserLeft(name string) string {
	return "Sayonara Kak " + name + "\n" +
		"Jangan lupa sama squad Bandung yaa 👋🏻👋🏻👋🏻"
}

func OnboardingUser(firstName, lastName string) string {
	return "Kak " + firstName + lastName + ", yuk perkenalan dulu. Tolong tuliskan biodata dengan format seperti ini\n\n" +
		"🐾 Nama panggilan : \n" +
		"🐾 Job title : \n" +
		"🐾 Squad : \n" +
		"🐾 Pekerjaan atau Pendidikan terakhir : \n" +
		"🐾 Status : \n" +
		"🐾 Hobi : \n" +
		"🐾 Motto : \n\n" +
		"*** Sebaiknya gunakan <b>HURUF</b> dan <b>ANGKA</b> saja.\n" +
		"Jika ada kendala harap hubungi Kak @mpermperpisang yaa #semangat"
}

func ValidOnboarding() string {
	return "Langkah selanjutnya\n" +
		"1. Isi biodata Kakak di tiny.cc/bukabandung\n" +
		"2. Subscribe official channel https://t.me/joinchat/AAAAAFScH6zPovO-LR_9nQ\n" +
		"3. Gabung salah 1 dari grup-grup berikut :\n\n" +
		"🐾 https://t.me/joinchat/AJmnGxEJVPlDllU198GfIg Grup Kuliner\n" +
		"🐾 https://t.me/joinchat/C3VkcA74oY2MUAL9_uVjVw Grup Badminton\n" +
		"🐾 https://t.me/joinchat/HgqJSBGPu9Rl-OpceBHwUw Grup Action figure\n" +
		"🐾 https://t.me/joinchat/EH0nV0kmrExNs7dmMD4fAA Grup Patungan kado\n" +
		"🐾 Grup Futsal. Silahkan menghubungi Kak @ivantedja\n\n" +
		"Selamat bergabung dengan Bukalapak Bandung yaa, Kak. Let's rock and roll 🎉🥳🤟🏻"
}

func InvalidOnboarding() string {
	return "Pengisian biodatanya masih salah. Ingat, Kak.\n" +
		"*** Sebaiknya gunakan <b>HURUF</b> dan <b>ANGKA</b> saja.\n" +
		"*** Chat ini tidak akan hilang sebelum Kakak mengirimkan biodata dengan format yang benar 😈\n" +
		"Jika ada kendala harap hubungi Kak @mpermperpisang yaa #semangat"
}

func OnboardingSuccess(text, username string) string {
	return "Berikut perkenalan dari Kak @" + username + "\n\n" + text +
		"\n\nTolong disambit #eh disambut dengan meriah yaa Kakaknya 🎉🥳🤟🏻"
=======
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
>>>>>>> 0bf34c0d62a747c009237eed2cf0ede2388788f7
}
