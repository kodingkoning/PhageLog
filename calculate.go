package main

import (
	"strconv"
)

func findRank(n int) string {
	if n%10 == 1 && n != 11 {
		return strconv.Itoa(n) + "st"
	} else if n%10 == 2 && n != 12 {
		return strconv.Itoa(n) + "nd"
	} else if n%10 == 3 && n != 13 {
		return strconv.Itoa(n) + "rd"
	} else {
		return strconv.Itoa(n) + "th"
	}
}
