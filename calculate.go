/* calculate.go provides calculations and helper methods for the phageLog project
 * Author: Elizabeth Koning
 * Written in 2017
 */
package main

import (
	"strconv"
)

// find rank converts an integer to a "rank"
// Parameter: n, an int
// Return: a string represenation such as "1st", "2nd", or "3rd"
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
