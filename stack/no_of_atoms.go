package main

import (
	"fmt"
	"maps"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func countOfAtoms(formula string) string {
	stack := make([]map[string]int, 0)
	i := 0
	stack = append(stack, map[string]int{})

	for i < len(formula) {
		if formula[i] == '(' {
			stack = append(stack, map[string]int{})
			i++
			continue
		}

		if formula[i] == ')' {
			current := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			i++

			if i < len(formula) && formula[i] >= '1' && formula[i] <= '9' {
				num, idx := getDigit(formula, i)

				for k := range current {
					current[k] *= num
				}

				i = idx
			}

			last := stack[len(stack)-1]
			for k := range current {
				last[k] += current[k]
			}

			stack[len(stack)-1] = last
			continue
		}

		current := stack[len(stack)-1]
		var str strings.Builder
		str.WriteByte(formula[i])

		i++

		for i < len(formula) && formula[i] >= 'a' && formula[i] <= 'z' {
			str.WriteByte(formula[i])
			i++
		}

		ele := str.String()

		if _, ok := current[ele]; !ok {
			current[ele] = 0
		}

		if i < len(formula) && formula[i] >= '1' && formula[i] <= '9' {
			num, idx := getDigit(formula, i)
			current[ele] += num
			i = idx
		} else {
			current[ele] += 1
		}

		stack[len(stack)-1] = current
	}

	resultMap := stack[0]
	var result strings.Builder
	keys := slices.Sorted(maps.Keys(resultMap))

	for _, key := range keys {
		if resultMap[key] == 1 {
			fmt.Fprintf(&result, "%s", key)
		} else {
			fmt.Fprintf(&result, "%s%d", key, resultMap[key])
		}
	}

	return result.String()
}

func getDigit(formula string, i int) (int, int) {
	var str strings.Builder
	for i < len(formula) && formula[i] >= '0' && formula[i] <= '9' {
		str.WriteByte(formula[i])
		i++
	}

	num, _ := strconv.Atoi(str.String())

	return num, i
}

func countOfAtomsUsingRegexp(formula string) string {
	// Regular expression to extract atom, count, (, ), multiplier
	// Each match has 5 capture groups (indices 1..5)
	//
	// Example: formula = "K4(ON(SO3)2)2"
	// FindAllStringSubmatch scans left to right and, at each position, matches
	// exactly one of the three alternatives (atom | "(" | ")"+digits), leaving
	// the other groups as "":
	//   "K4" -> group1="K" group2="4"        (atom "K", count 4)
	//   "("  -> group3="("                    (open a new scope)
	//   "O"  -> group1="O"                    (atom "O", implicit count 1)
	//   "N"  -> group1="N"                    (atom "N", implicit count 1)
	//   "("  -> group3="("                    (nested scope)
	//   "S"  -> group1="S"                    (atom "S", implicit count 1)
	//   "O3" -> group1="O" group2="3"         (atom "O", count 3)
	//   ")2" -> group4=")" group5="2"         (close scope, multiply counts by 2)
	//   ")2" -> group4=")" group5="2"         (close outer scope, multiply by 2)
	re := regexp.MustCompile(`([A-Z][a-z]*)(\d*)|(\()|(\))(\d*)`)
	matches := re.FindAllStringSubmatch(formula, -1)

	// Stack to keep track of the atoms and their counts
	stack := []map[string]int{{}}

	for _, m := range matches {
		atom, count, left, right, multiplier := m[1], m[2], m[3], m[4], m[5]

		switch {
		// If atom, add it to the top hashmap
		case atom != "":
			c := 1
			if count != "" {
				c, _ = strconv.Atoi(count)
			}
			stack[len(stack)-1][atom] += c

		// If left parenthesis, push a new hashmap onto the stack
		case left != "":
			stack = append(stack, map[string]int{})

		// If right parenthesis, pop the top map, multiply its counts,
		// and merge them into the new top map
		case right != "":
			currMap := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			mult := 1
			if multiplier != "" {
				mult, _ = strconv.Atoi(multiplier)
			}

			top := stack[len(stack)-1]
			for a, c := range currMap {
				top[a] += c * mult
			}
		}
	}

	// Sort the atom names
	finalMap := stack[0]
	atoms := make([]string, 0, len(finalMap))
	for a := range finalMap {
		atoms = append(atoms, a)
	}
	sort.Strings(atoms)

	// Generate the answer string
	var sb strings.Builder
	for _, a := range atoms {
		sb.WriteString(a)
		if finalMap[a] > 1 {
			sb.WriteString(strconv.Itoa(finalMap[a]))
		}
	}

	return sb.String()
}

func preProcessingAtomCount(formula string) []int {
	stack := make([]int, 0)
	mul := 1
	var currCount strings.Builder
	mulFactor := make([]int, len(formula))

	for i := len(formula) - 1; i >= 0; i-- {
		if formula[i] == ')' {
			digit := 1
			if currCount.String() != "" {
				digitString := reverse(currCount.String())
				digit, _ = strconv.Atoi(digitString)
			}

			mul *= digit
			stack = append(stack, digit)
			currCount.Reset()

			continue
		}

		if formula[i] == '(' {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			mul /= top
			continue
		}

		if formula[i] >= '0' && formula[i] <= '9' {
			currCount.WriteByte(formula[i])
		} else {
			currCount.Reset()
		}

		mulFactor[i] = mul
	}

	return mulFactor
}

func countOfAtomsUsingPreProcessing(formula string) string {
	mulFactor := preProcessingAtomCount(formula)
	resultMap := make(map[string]int)
	i := 0

	for i < len(formula) {
		if formula[i] < 'A' || formula[i] > 'Z' {
			i++
			continue
		}

		var str strings.Builder
		str.WriteByte(formula[i])
		mul := mulFactor[i]

		i++

		for i < len(formula) && formula[i] >= 'a' && formula[i] <= 'z' {
			str.WriteByte(formula[i])
			i++
		}

		ele := str.String()

		if i < len(formula) && formula[i] >= '1' && formula[i] <= '9' {
			num, idx := getDigit(formula, i)
			resultMap[ele] += num * mul
			i = idx
		} else {
			resultMap[ele] += mul
		}
	}

	var result strings.Builder
	keys := slices.Sorted(maps.Keys(resultMap))

	for _, key := range keys {
		if resultMap[key] == 1 {
			fmt.Fprintf(&result, "%s", key)
		} else {
			fmt.Fprintf(&result, "%s%d", key, resultMap[key])
		}
	}

	return result.String()
}

func reverse(s string) string {
	b := []byte(s)
	slices.Reverse(b)
	return string(b)
}

func countOfAtomsUsingPreProcessingSimplified(formula string) string {
	n := len(formula)

	// Pass 1 (right to left): compute the multiplier that applies at each position.
	mulFactor := make([]int, n)
	stack := []int{}
	mul, num, place := 1, 0, 1

	for i := n - 1; i >= 0; i-- {
		c := rune(formula[i])
		switch {
		case unicode.IsDigit(c):
			num += int(c-'0') * place // build the number without reversing
			place *= 10
			continue // keep accumulating; don't reset num
		case c == ')':
			m := max(num, 1)
			mul *= m
			stack = append(stack, m)
		case c == '(':
			mul /= stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		default: // a letter
			mulFactor[i] = mul
		}
		num, place = 0, 1 // any non-digit ends the current number
	}

	// Pass 2 (left to right): read each atom and its count.
	counts := map[string]int{}
	for i := 0; i < n; {
		if !unicode.IsUpper(rune(formula[i])) {
			i++
			continue
		}

		start := i
		i++
		for i < n && unicode.IsLower(rune(formula[i])) {
			i++
		}
		atom := formula[start:i]

		cnt := 0
		for i < n && unicode.IsDigit(rune(formula[i])) {
			cnt = cnt*10 + int(formula[i]-'0')
			i++
		}

		counts[atom] += max(cnt, 1) * mulFactor[start]
	}

	// Build the result in sorted order.
	var sb strings.Builder
	for _, atom := range slices.Sorted(maps.Keys(counts)) {
		sb.WriteString(atom)
		if counts[atom] > 1 {
			sb.WriteString(strconv.Itoa(counts[atom]))
		}
	}
	return sb.String()
}

func noOfAtomsMain() {
	formula := "K4(ON(SeO3)2)2"
	fmt.Println(countOfAtoms(formula))
	fmt.Println(countOfAtomsUsingRegexp(formula))
	fmt.Println(countOfAtomsUsingPreProcessing(formula))
	fmt.Println(countOfAtomsUsingPreProcessingSimplified(formula))

	formula = "H11He49NO35B7N46Li20"
	fmt.Println(countOfAtoms(formula))
	fmt.Println(countOfAtomsUsingRegexp(formula))
	fmt.Println(countOfAtomsUsingPreProcessing(formula))
	fmt.Println(countOfAtomsUsingPreProcessingSimplified(formula))
}
