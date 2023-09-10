package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	//-----------------------------------
	fu_sliding_window_counter()
	//-----------------------------------
	generator_receiver()
	//-----------------------------------
	odd_even_print()
	//-----------------------------------
	ticker()
	//-----------------------------------
	what_output()
	//-----------------------------------
	worker_pool_extended()
	//-----------------------------------
	worker_pool()
}
