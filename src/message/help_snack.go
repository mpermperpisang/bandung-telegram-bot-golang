package message

import (
	"os"
)

func Snack() string {
	var botName = os.Getenv("BOT_SNACK")

	return "<b>DAY FORMAT</b>\n" +
		"only use <code>mon</code>, <code>tue</code>, <code>wed</code>, <code>thu</code> or <code>fri</code>\n\n" +

		"1. Add people to snack schedule (GROUP)\n" +
		"<b>/add_snack@" + botName + " day1 @username day2 @username day3 @username</b>\n\n" +

		"2. Change people schedule TEMPORARILY (GROUP)\n" +
		"<b>/move@" + botName + " day1 @username day2 @username day3 @username</b>\n\n" +

		"3. Change snack schedule PERMANENTLY (GROUP)\n" +
		"<b>/permanent@" + botName + " day1 @username day2 @username day3 @username</b>\n\n" +

		"4. Delete people from schedule (GROUP)\n" +
		"<b>/delete@" + botName + " @username1 @username2 @username3</b>\n\n" +

		"5. Cancel people from bring snack (GROUP)\n" +
		"<b>/cancel@" + botName + " @username1 @username2 @username3</b>\n\n" +

		"6. People brought the snack (PRIVATE)\n" +
		"<b>/done@" + botName + "</b>\n\n" +

		"7. Free people from snack schedule (GROUP)\n" +
		"<b>/holiday@" + botName + " @all</b>\n" +
		"or\n" +
		"<b>/holiday@" + botName + " @username1 @username2 @username3</b>\n\n" +

		"8. Send dear all message to all of Bukalapak.bdg member (GROUP)\n" +
		"<b>/dear_all@" + botName + "</b>\n\n" +

		"9. Add admin snack (GROUP)\n" +
		"<b>/add_admin@" + botName + " @username1 @username2 @username3</b>\n\n" +

		"10. See the list of admin snack (PRIVATE)\n" +
		"<b>/list_admin@" + botName + "</b>\n\n" +

		"11. See list of snack schedule (PRIVATE)\n" +
		"<b>/schedule@" + botName + "</b>\n\n" +

		"🐾 Only admin/managers can do add, edit, change, " +
		"delete & holiday 😎"
}
