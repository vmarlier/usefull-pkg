package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(formatNumber("06 79 13 68 44"))
	fmt.Println(formatNumber("06.23.10.10.52"))
	fmt.Println(formatNumber("+216 51819956"))
	fmt.Println(formatNumber("+33 0679136844"))
}

func formatNumber(num string) string {
	var number string
	var newNumber string

	r := []rune(num)

	for _, n := range r {
		if n >= 48 && n <= 57 || n == 43 {
			number += string(n)
		}
	}

	if strings.Contains(number, "+") {
		if strings.Contains(number, "+33") {
			r := []rune(number)

			for i, n := range r {
				if i != 3 && string(n) != "0" {
					newNumber += string(n)
				}
			}
			return newNumber
		}
	} else {
		r := []rune(number)

		for i, n := range r {
			if i == 0 && string(n) == "0" {
				newNumber += "+33"
			} else {
				newNumber += string(n)
			}
		}

		return newNumber
	}
	return number
}
