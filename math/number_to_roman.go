package main

import "fmt"

func Romanize(number int) string {
	lookup := []struct {
		Value  int
		Symbol string
	}{
		{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"}, {100, "C"},
		{90, "XC"}, {50, "L"}, {40, "XL"}, {10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
	}

	roman := ""

	for _, entry := range lookup {
		for number >= entry.Value {
			roman += entry.Symbol
			number -= entry.Value
		}
	}
	return roman
}

// numberToRoman converts a number to a Roman numeral
func numberToRoman() {
	number := 8
	roman := Romanize(number)
	fmt.Printf("Roman numeral for %d is %s\n", number, roman)
}
