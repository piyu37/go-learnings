package main

import "fmt"

func GenerateDivTags(numberOfTags int) []string {
	result := createDivtags("", numberOfTags, numberOfTags)
	return result
}

func createDivtags(partialString string, openingTag, closingTag int) []string {
	result := make([]string, 0)

	if closingTag == 0 {
		result = append(result, partialString)
		return result
	}

	if openingTag > 0 {
		result = append(result, createDivtags(partialString+"<div>", openingTag-1, closingTag)...)
	}

	if closingTag > openingTag {
		result = append(result, createDivtags(partialString+"</div>", openingTag, closingTag-1)...)
	}

	return result
}

// https://github.com/lee-hen/Algoexpert/tree/master/hard/32_generate_div_tags
func generateDivTagsMain() {
	fmt.Println(GenerateDivTags(3))
}
