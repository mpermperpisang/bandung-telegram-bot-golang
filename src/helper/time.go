package helper

import (
	"time"
)

var day, date, month, year string

func DayNow() string {
	t := time.Now()
	day = t.Format("Mon")

	return day
}

func DateNow() string {
	t := time.Now()
	date = t.Format("02")

	return date
}

func MonthNow() string {
	t := time.Now()
	month = t.Format("01")

	return month
}

func YearNow() string {
	t := time.Now()
	year = t.Format("2006")

	return year
}
