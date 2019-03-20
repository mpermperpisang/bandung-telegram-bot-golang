package database

func UpdateOnboarding(username string) {
	db := DBConnection()
	db.Exec("UPDATE onboarding_member SET on_flag='true' WHERE on_username='@" + username + "'")
}
