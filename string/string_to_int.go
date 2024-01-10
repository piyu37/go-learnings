package main

import (
	"fmt"
)

func stringToInt(s string) (int, error) {
	result := 0
	sign := 1
	for i, char := range s {
		if i == 0 && char == '-' {
			sign = -1
			continue
		}
		digit := int(char - '0')
		if digit < 0 || digit > 9 {
			return 0, fmt.Errorf("invalid character in the string: %c", char)
		}
		result = result*10 + digit
	}
	return result * sign, nil
}

func stringToIntMain() {
	input := "12345"
	intValue, err := stringToInt(input)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Converted value: %d\n", intValue)
	}
}
