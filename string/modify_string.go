package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
 * Complete the 'smallestString' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func smallestString(s string) string {
	for i := 0; i < len(s); i++ {
		if rune(s[i]) != 'c' {
			continue
		}

		if i+1 < len(s) && rune(s[i+1]) == 'b' {
			s = s[:i] + "bc" + s[i+2:]

			for j := i - 1; j >= 0; j-- {
				if rune(s[j]) != 'c' {
					continue
				}

				if j+1 < len(s) && rune(s[j+1]) == 'b' {
					s = s[:j] + "bc" + s[j+2:]
				}
			}
		}
	}

	for i := 0; i < len(s); i++ {
		if rune(s[i]) != 'b' {
			continue
		}

		if i+1 < len(s) && rune(s[i+1]) == 'a' {
			s = s[:i] + "ab" + s[i+2:]

			for j := i - 1; j >= 0; j-- {
				if rune(s[j]) != 'b' {
					continue
				}

				if j+1 < len(s) && rune(s[j+1]) == 'a' {
					s = s[:j] + "ab" + s[j+2:]
				}
			}
		}
	}

	return s
}

// modifyString, takes a string s as input and modifies it to produce the smallest lexicographically
// possible string with the following rules:

// If there is a 'c' followed by 'b', it replaces 'cb' with 'bc'.
// This process is repeated until no more 'cb' is found.
// If there is a 'b' followed by 'a', it replaces 'ba' with 'ab'.
// This process is repeated until no more 'ba' is found.
func modifyString() {
	// reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, _ := os.Create(os.Getenv("OUTPUT_PATH"))
	// checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := "acbba"

	result := smallestString(s)

	fmt.Fprintf(writer, "%s\n", result)

	writer.Flush()
}
