package message

func Header() string {
	return "Yohoo,,,\n<b>PLEASE READ AND REMEMBER</b>\n\n"
}

func Footer() string {
	return "\n\n👣 <b>GROUP</b> means you can only send these commands in your Telegram group\n" +
		"👣 <b>PRIVATE</b> means you can only send these commands if you're chat with me\n" +
		"👣 <b>BOTH</b> means you can send these commands either in group or private chat with me\n\n" +
		"<code>Invalid command will give a response to private message. So, make sure you're not blocked the bot and click START</code>\n\n" +
		"want to help improve this bot? or report the bugs?\nprivate message @mpermperpisang, please ☺️"
}
