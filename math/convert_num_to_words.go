package main

import (
	"fmt"
)

// Strings at index 0 are not used, they are to make array indexing simple
var one = []string{"", "one ", "two ", "three ", "four ",
	"five ", "six ", "seven ", "eight ",
	"nine ", "ten ", "eleven ", "twelve ",
	"thirteen ", "fourteen ", "fifteen ",
	"sixteen ", "seventeen ", "eighteen ",
	"nineteen "}

// Strings at index 0 and 1 are not used, they are to make array indexing simple
var ten = []string{"", "", "twenty ", "thirty ", "forty ",
	"fifty ", "sixty ", "seventy ", "eighty ",
	"ninety "}

// NumToWords converts a 1- or 2-digit number to words
func NumToWords(n int, s string) string {
	var str string

	// If n is more than 19, divide it
	if n > 19 {
		str += ten[n/10] + one[n%10]
	} else {
		str += one[n]
	}

	// If n is non-zero
	if n != 0 {
		str += s
	}

	return str
}

// ConvertToWordsUpto9digitAndIndianCount converts a given number to words
// Long handles up to 9-digit number, change to unsigned long long
// int to handle more digit number
func ConvertUpto9DigitNumToWordsWithIndianCount(n int) string {
	var out string

	// Handles digits at ten millions and hundred millions places (if any)
	out += NumToWords(n/10000000, "crore ")

	// Handles digits at hundred thousands and one millions places (if any)
	out += NumToWords((n/100000)%100, "lakh ")

	// Handles digits at thousands and tens thousands places (if any)
	out += NumToWords((n/1000)%100, "thousand ")

	// Handles digit at hundreds places (if any)
	out += NumToWords((n/100)%10, "hundred ")

	if n > 100 && n%100 != 0 {
		out += "and "
	}

	// Handles digits at ones and tens places (if any)
	out += NumToWords(n%100, "")

	return out
}

// ConvertNumToWordsWithUSCount converts a given number to words
// Long handles up to 9-digit number, change to unsigned long long
// int to handle more digit number
func ConvertNumToWordsWithUSCount(n int) string {
	limit, t := 1000000000000, 0

	// If zero, print zero
	if n == 0 {
		return "zero"
	}

	// Array to store the powers of 10
	multiplier := []string{"", "trillion", "billion", "million", "thousand"}

	// Array to store multiples of ten
	tens := []string{"", "twenty", "thirty", "forty", "fifty",
		"sixty", "seventy", "eighty", "ninety"}

	// If number is less than 20, print without any
	if n < 20 {
		return one[n]
	}

	var answer string
	i := n
	for i > 0 {
		// Store the value in multiplier[t], i.e., n = 1000000
		// Then r = 1 for multiplier (million), 0 for multipliers (trillion and billion)
		// Multiplier here refers to the current accessible limit
		currHun := i / limit

		// It might be possible that the current multiplier is bigger than your number
		for currHun == 0 {
			// Set i as the remainder obtained when n was divided by the limit
			i %= limit

			// Divide the limit by 1000, shifts the multiplier
			limit /= 1000

			// Get the current value in hundreds, as the English system works in hundreds
			currHun = i / limit

			// Shift the multiplier
			t++
		}

		// If current hundred is greater than 99, add the hundreds' place
		if currHun > 99 {
			answer += (one[currHun/100] + " hundered ")
		}

		// Bring the current hundred to tens
		currHun = currHun % 100

		// If the value in tens belongs to [1,19], add using the one
		if currHun > 0 && currHun < 20 {
			answer += (one[currHun] + " ")
		} else if currHun%10 == 0 && currHun != 0 {
			// If currHun is now a multiple of 10 but not 0
			// Add the tens' value using the tens array
			answer += (tens[(currHun/10)-1] + " ")
		} else if currHun > 19 && currHun < 100 {
			// If the value belongs to [21,99], excluding the multiples of 10
			// Get the ten's place and one's place, and print using the one array
			answer += (tens[(currHun/10)-1] + " " +
				one[currHun%10] + " ")
		}

		// If Multiplier has not become less than 1000, shift it
		if t < 4 {
			t++
			answer += (multiplier[t] + " ")
		}

		i = i % limit
		limit = limit / 1000
	}

	return answer
}

func convertNumToWords() {
	n := 438237764

	// Convert given number to words
	fmt.Println(ConvertUpto9DigitNumToWordsWithIndianCount(n))

	n = 999999999

	// Convert given number to words
	fmt.Println(ConvertUpto9DigitNumToWordsWithIndianCount(n))

	n1 := 10101010110001

	// Convert given number to words
	fmt.Println(ConvertNumToWordsWithUSCount(n1))
}
