package main

import "fmt"

func myfun(a interface{}) {
	value, ok := a.(float64)
	fmt.Println(value, ok)
}

func interfaceTypeModelling() {
	var a1 interface{} = 99.09
	myfun(a1) // 99.09 true

	var a2 interface{} = "WashingPowder"
	myfun(a2) // 0 false
}

func myfunc2(a interface{}) {
	val := a.(int)
	fmt.Println("Value:", val)
}

func interfaceTypeModelling2() {
	var val interface{} = "WashingPowder"
	myfunc2(val) // panic: interface conversion: interface {} is string, not int
}
