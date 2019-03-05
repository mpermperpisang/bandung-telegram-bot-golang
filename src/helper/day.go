package helper

func DayName(day string) string {
	var dayName string

	switch day {
	case "mon":
		dayName = "Senin"
	case "tue":
		dayName = "Selasa"
	case "wed":
		dayName = "Rabu"
	case "thu":
		dayName = "Kamis"
	case "fri":
		dayName = "Jumat"
	default:
		dayName = "Libur"
	}

	return dayName
}
