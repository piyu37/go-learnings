package main

import "fmt"

var phoneNemonicsMap map[string][]string = map[string][]string{
	"1": {"1"},
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
	"0": {"0"},
}

func PhoneNumberMnemonics(phoneNumber string) []string {
	result := []string{}

	for _, ru := range phoneNumber {
		char := string(ru)

		if len(result) == 0 {
			allChars := append([]string{}, phoneNemonicsMap[char]...)

			result = append(result, allChars...)

			continue
		}

		tempResult := append([]string{}, result...)
		result = []string{}

		for v := range tempResult {
			allChars := append([]string{}, phoneNemonicsMap[char]...)

			for i := range allChars {
				ch := allChars[i]

				result = append(result, tempResult[v]+ch)
			}
		}
	}

	return result
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/41_phone_number_mnemonics
func mnemonics() {
	phoneNum := "1905"

	fmt.Println(PhoneNumberMnemonics(phoneNum))
}
