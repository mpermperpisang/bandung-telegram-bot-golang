package message

import "os"

func Header() string {
	return "Yohoo,,,\n<b>PLEASE READ AND REMEMBER</b>\n\n"
}

func Booking() string {
	return "1. Requesting deploy a branch (GROUP)\n" +
		"<b>/deploy_request@" + os.Getenv("BOT_BOOKING") + " branch_name</b>\n\n" +

		"2. Cancel deploy branch request (GROUP)\n" +
		"<b>/cancel_request@" + os.Getenv("BOT_BOOKING") + " branch_name</b>\n\n" +

		"3. Booking staging for hard code, deploy or testing (GROUP)\n" +
		"<b>/booking@" + os.Getenv("BOT_BOOKING") + " staging_number</b>\n\n" +

		"4. Done using staging (GROUP)\n" +
		"<b>/done@" + os.Getenv("BOT_BOOKING") + " staging_number</b>\n\n" +

		"5. Add staging to squad (GROUP)\n" +
		"<b>/add_staging@" + os.Getenv("BOT_BOOKING") + " squad_name staging_number1 staging_number2 staging_number3</b>\n\n" +

		"6. See staging book status (GROUP)\n" +
		"<b>/status@" + os.Getenv("BOT_BOOKING") + " staging_number1 staging_number2 staging_number3</b>\n" +
		"or\n" +
		"<b>/status@" + os.Getenv("BOT_BOOKING") + " squad_name</b>\n\n" +

		"7. See list of deploy branch request (GROUP)\n" +
		"<b>/list_request@" + os.Getenv("BOT_BOOKING") + "</b>\n\n" +

		"🐾 You can use another staging, not only DANA staging"
}

func Footer() string {
	return "\n\n<b>GROUP</b> means you can only send these commands in your Telegram group\n" +
		"<b>PRIVATE</b> means you can only send these commands if you chat with me\n" +
		"<b>BOTH</b> means you can send these commands either in group or private chat with me\n\n" +
		"want to help improve this bot?\nprivate message @mpermperpisang, please ☺️"
}
