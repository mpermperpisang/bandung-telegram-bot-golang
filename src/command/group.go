package command

func Group() {
	GBooking()
	GSnack()
}

func GBooking() {
	// List command group untuk bot Booking
	MatchStatusStaging()
	MatchAddStaging()
	MatchUpdateStaging()
	MatchBookStaging()
	MatchDoneStaging()
	MatchAddOnCall()
	MatchOnCall()
}

func GSnack() {
	// List command group untuk bot Snack
	MatchAddSnack()
	MatchMoveSnack()
	MatchPermanentSnack()
	MatchDeleteSnack()
	MatchCancelSnack()
	MatchHolidaySnack()
	MatchAddAdminSnack()
	MatchDeleteAdminSnack()
	MatchListAdminSnack()
	MatchBaikSnack()
	MatchAskSnack()
}
