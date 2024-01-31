package main

func NumberOfWaysToTraverseGraph2(width int, height int) int {
	xDistanceToCorner := width - 1
	yDistanceToCorner := height - 1

	numerator := factorial(xDistanceToCorner + yDistanceToCorner)
	denominator := factorial(xDistanceToCorner) * factorial(yDistanceToCorner)
	return numerator / denominator
}

func factorial(num int) int {
	result := 1
	for n := 2; n < num+1; n++ {
		result *= n
	}
	return result
}
