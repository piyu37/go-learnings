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

// Deromanize converts a Roman numeral to its integer value.
// A symbol smaller than the one after it is subtracted (e.g. IV = 4), otherwise added.
func Deromanize(roman string) int {
	values := map[byte]int{
		'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
	}

	number := 0

	for i := 0; i < len(roman); i++ {
		current := values[roman[i]]
		if i+1 < len(roman) && current < values[roman[i+1]] {
			number -= current
		} else {
			number += current
		}
	}
	return number
}

// numberToRoman converts a number to a Roman numeral
func numberToRoman() {
	number := 8
	roman := Romanize(number)
	fmt.Printf("Roman numeral for %d is %s\n", number, roman)

	romNum := "VIII"
	fmt.Println(Deromanize(romNum))
}
