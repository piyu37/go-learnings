package main

import (
	"fmt"
	"strings"
)

func PhoneNumberMnemonicsRecur(phoneNumber string) []string {
	currentMnemonic := make([]string, len(phoneNumber))

	for i := range currentMnemonic {
		currentMnemonic[i] = "0"
	}

	result := make([]string, 0)

	phoneNumMnemonicHelper(0, phoneNumber, currentMnemonic, &result)

	return result
}

func phoneNumMnemonicHelper(idx int, phoneNumber string, currentMnemonic []string, result *[]string) {
	if idx == len(phoneNumber) {
		mnemonic := strings.Join(currentMnemonic, "")
		*result = append(*result, mnemonic)

		return
	}

	ch := string(phoneNumber[idx])

	allChars := phoneNemonicsMap[ch]

	for i := range allChars {
		c := allChars[i]

		currentMnemonic[idx] = c

		phoneNumMnemonicHelper(idx+1, phoneNumber, currentMnemonic, result)
	}
}

// https://github.com/lee-hen/Algoexpert/tree/master/medium/41_phone_number_mnemonics
func mnemonicsRecur() {
	phoneNum := "1905"

	fmt.Println(PhoneNumberMnemonicsRecur(phoneNum))
}
