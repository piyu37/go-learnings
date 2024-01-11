package main

import "fmt"

func ShortenPath(path string) string {
	result := ""

	str := ""
	stack := make([]string, 0)

	for _, ch := range path {
		str += string(ch)

		if string(ch) == "/" {
			if len(stack) > 0 {
				if (str == "/") || (str == "./") || (str == "../" && stack[len(stack)-1] == "/") {
					str = ""

					continue
				} else if str == "../" && stack[len(stack)-1] != "../" {
					stack = stack[:len(stack)-1]
				} else {
					stack = append(stack, str)
				}
			} else {
				if str == "./" {
				} else {
					stack = append(stack, str)
				}
			}

			str = ""
		}
	}

	if len(stack) > 0 {
		if (str == "/") || (str == "./") || (str == ".." && stack[len(stack)-1] == "/") {
		} else if str == ".." && stack[len(stack)-1] != "../" {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, str)
		}
	} else {
		if str == "./" {
			stack = append(stack, "")
		} else {
			stack = append(stack, str)
		}
	}

	for _, val := range stack {
		result += val
	}

	if len(result) > 0 && string(result[len(result)-1]) == "/" {
		result = result[:len(result)-1]
	}

	return result
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/41_shorten_path
func shortenPathMain() {
	fmt.Println(ShortenPath("./.."))
}
