package nil

import (
	"strconv"
	"time"
)

const Version string = "\033[1;35mV S-1.0.0"

func GenerateID() string {
	t := time.Now()
	year, month, day := t.Date()
	formatMonth := "SET"
	switch month {
	case time.January:
		formatMonth = "JAN"
	case time.February:
		formatMonth = "FEB"
	case time.March:
		formatMonth = "MAR"
	case time.April:
		formatMonth = "APR"
	case time.May:
		formatMonth = "MAY"
	case time.June:
		formatMonth = "JUN"
	case time.July:
		formatMonth = "JUL"
	case time.August:
		formatMonth = "AUG"
	case time.September:
		formatMonth = "SEP"
	case time.October:
		formatMonth = "OCT"
	case time.November:
		formatMonth = "NOV"
	case time.December:
		formatMonth = "DEC"
	default:
		formatMonth = "NIL"
	}
	hour, min, _ := t.Clock()

	strday := strconv.Itoa(day)
	stryear := strconv.Itoa(year)
	strhour := strconv.Itoa(hour)
	strmin := strconv.Itoa(min)

	uuid := strday + "-" + formatMonth + "-" + stryear + "~" + strhour + ":" + strmin

	return uuid
}
