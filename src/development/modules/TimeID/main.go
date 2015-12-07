package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	month := t.Month()
	formatMonth := 0
	switch month {
	case "January":
		formatMonth = 1
	case "February":
		formatMonth = 2
	case "March":
		formatMonth = 3
	case "April":
		formatMonth = 4
	case "May":
		formatMonth = 5
	case "June":
		formatMonth = 6
	case "July":
		formatMonth = 7
	case "August":
		formatMonth = 8
	case "September":
		formatMonth = 9
	case "October":
		formatMonth = 10
	case "November":
		formatMonth = 11
	case "December":
		formatMonth = 12
	default:
		formatMonth = 0
	}
	hour := t.Hour()
}
