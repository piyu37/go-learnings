package main

import (
	"fmt"
	"math/rand"
)

var lowerChars = []string{
	"a", "b", "c",
}

var specialChars = []string{
	"!", "@", "#",
}

var numChars = []string{
	"1", "2", "3",
}

func randInt(min int, max int) int {
	return min + rand.Intn(max-min)
}

func giveRandomFromUpper() string {
	v := byte(randInt(65, 90))
	return string(v)
}

func giveRandomFromLower() string {
	v := byte(randInt(97, 123))
	return string(v)
}

func giveRandomFromSpecial() string {
	return specialChars[rand.Intn(len(specialChars))]
}

func giveFourChar() {

}

func giveRandomFromNumber() string {
	v := byte(randInt(55, 64))
	return string(v)
}

func generatePassowrd() string {
	return ""
}

// Program to generate 5 random password contain 1 upper case, 1 lower case, 1 special, 1 no., len=8
func generateRandomPasswords() {
	for i := 0; i < 5; i++ {
		password := generatePassowrd()

		fmt.Println(password)
	}
}
