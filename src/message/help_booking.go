package message

import (
	"os"
)

func Booking() string {
	var botName = os.Getenv("BOT_BOOKING")

	return "1. Add staging to squad (GROUP)\n" +
		"<b>/add_staging@" + botName + " squad_name staging_number1 staging_number2 staging_number3</b>\n\n" +

		"2. Update staging move to another squad (GROUP)\n" +
		"<b>/update_staging@" + botName + " squad_name staging_number1 staging_number2 staging_number3</b>\n\n" +

		"3. Booking staging for hard code, deploy or testing (GROUP)\n" +
		"<b>/booking@" + botName + " staging_number</b>\n\n" +

		"4. Done using staging (GROUP)\n" +
		"<b>/done@" + botName + " staging_number</b>\n\n" +

		"6. See staging book status (GROUP)\n" +
		"<b>/status_staging@" + botName + " staging_number1 staging_number2 staging_number3</b>\n" +
		"or\n" +
		"<b>/status_staging@" + botName + " squad_name</b>\n\n" +

		"🐾 You can use another staging"
}