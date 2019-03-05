package command

func Private() {
	PGeneral()
	PSnack()
}

func PGeneral() {
	// List command private untuk semua bot
	MatchStart()
	MatchHelp()
}

func PSnack() {
	// List command private untuk bot Snack
	MatchDoneSnack()
}
