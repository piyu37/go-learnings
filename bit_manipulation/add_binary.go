package main

import (
	"fmt"
	"math/big"
	"slices"
)

func addBinary(a string, b string) string {
	// Convert binary strings to big.Int
	numA := new(big.Int)
	numA, ok := numA.SetString(a, 2)
	if !ok {
		return "0"
	}

	numB := new(big.Int)
	numB, ok = numB.SetString(b, 2)
	if !ok {
		return "0"
	}

	// Sum the big.Int numbers
	sum := new(big.Int).Add(numA, numB)

	// Convert the sum back to a binary string
	return sum.Text(2)
}

func addBinary2(A string, B string) string {
	// The rules are:
	// a + b + carry
	// 0 + 0 + 0    = 0, carry = 0
	// 1 + 0 + 0    = 1, carry = 0
	// 1 + 1 + 0    = 0, carry = 1
	// 1 + 1 + 1    = 1, carry = 1
	var r []rune
	a := []rune(A)
	b := []rune(B)

	slices.Reverse(a)
	slices.Reverse(b)
	carry := 0
	for i := 0; i < len(a) || i < len(b); i++ {
		sum := carry
		if i < len(a) && a[i] == '1' {
			sum++
		}
		if i < len(b) && b[i] == '1' {
			sum++
		}
		switch sum {
		case 0:
			r = append(r, '0')
			carry = 0
		case 1:
			r = append(r, '1')
			carry = 0
		case 2:
			r = append(r, '0')
			carry = 1
		case 3:
			r = append(r, '1')
			carry = 1
		}
	}

	if carry == 1 {
		r = append(r, '1')
	}

	slices.Reverse(r)

	return string(r)
}

func addBinary3(a string, b string) string {
	var x, y, zero, carry, answer big.Int
	_, _ = x.SetString(a, 2)
	_, _ = y.SetString(b, 2)
	zero.SetInt64(0)
	for y.Cmp(&zero) != 0 {
		answer.Xor(&x, &y)
		carry.And(&x, &y)
		carry.Lsh(&carry, 1)
		x.Set(&answer)
		y.Set(&carry)
	}
	return x.Text(2)
}

// https://leetcode.com/problems/add-binary/?envType=study-plan-v2&envId=top-interview-150
func addBinaryMain() {
	a := "11"
	b := "1"

	fmt.Println(addBinary(a, b))
	fmt.Println(addBinary2(a, b))
	fmt.Println(addBinary3(a, b))
}
