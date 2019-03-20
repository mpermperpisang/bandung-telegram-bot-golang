package message

import (
	"os"
)

func Snack() string {
	var botName = os.Getenv("BOT_SNACK")

	return "<b>DAY FORMAT</b>\n" +
		"only use <code>mon</code>, <code>tue</code>, <code>wed</code>, <code>thu</code> or <code>fri</code>\n\n" +

		"1. Add people to snack schedule (BOTH)\n" +
		"<b>/add_snack@" + botName + " day1 @username day2 @username day3 @username</b>\n\n" +

		"2. Change people schedule <code>TEMPORARILY</code> (BOTH)\n" +
		"<b>/move@" + botName + " day1 @username day2 @username day3 @username</b>\n\n" +

		"3. Change snack schedule <code>PERMANENTLY</code> (BOTH)\n" +
		"<b>/permanent@" + botName + " day1 @username day2 @username day3 @username</b>\n\n" +

		"4. Delete people from schedule (BOTH)\n" +
		"<b>/delete@" + botName + " @username1 @username2 @username3</b>\n\n" +

		"5. Cancel people from bring snack (BOTH)\n" +
		"<b>/cancel@" + botName + " @username1 @username2 @username3</b>\n\n" +

		"6. People brought the snack (PRIVATE)\n" +
		"<b>/done@" + botName + "</b>\n\n" +

		"7. Free people from snack schedule (BOTH)\n" +
		"<b>/holiday@" + botName + " @all</b>\n" +
		"or\n" +
		"<b>/holiday@" + botName + " @username1 @username2 @username3</b>\n\n" +

		"8. Add admin snack (BOTH)\n" +
		"<b>/add_admin@" + botName + " @username1 @username2 @username3</b>\n\n" +

		"9. Delete admin snack (BOTH)\n" +
		"<b>/delete_admin@" + botName + " @username1 @username2 @username3</b>\n\n" +

		"10. See the list of admin snack (BOTH)\n" +
		"<b>/list_admin@" + botName + "</b>\n\n" +

		"11. See the kindest person (BOTH)\n" +
		"<b>/baik@" + botName + "</b>\n\n" +

		"🐾 Only admins can do add, edit, change, " +
		"delete & holiday 😎"
}
