package message

import (
	"os"
)

func Booking() string {
	var botName = os.Getenv("BOT_BOOKING")

	return "1. Requesting deploy a branch (GROUP)\n" +
		"<b>/deploy_request@" + botName + " branch_name</b>\n\n" +

		"2. Cancel deploy branch request (GROUP)\n" +
		"<b>/cancel_request@" + botName + " branch_name</b>\n\n" +

		"3. Booking staging for hard code, deploy or testing (GROUP)\n" +
		"<b>/booking@" + botName + " staging_number</b>\n\n" +

		"4. Done using staging (GROUP)\n" +
		"<b>/done@" + botName + " staging_number</b>\n\n" +

		"5. Add staging to squad (GROUP)\n" +
		"<b>/add_staging@" + botName + " squad_name staging_number1 staging_number2 staging_number3</b>\n\n" +

		"6. See staging book status (GROUP)\n" +
<<<<<<< HEAD
		"<b>/status_staging@" + botName + " staging_number1 staging_number2 staging_number3</b>\n" +
		"or\n" +
		"<b>/status_staging@" + botName + " squad_name</b>\n\n" +
=======
		"<b>/status@" + botName + " staging_number1 staging_number2 staging_number3</b>\n" +
		"or\n" +
		"<b>/status@" + botName + " squad_name</b>\n\n" +
>>>>>>> 94bcef742cbe47c008f1b4ba6cd7132bb3b99c64

		"7. See list of deploy branch request (GROUP)\n" +
		"<b>/list_request@" + botName + "</b>\n\n" +

		"🐾 You can use another staging, not only DANA staging"
}
