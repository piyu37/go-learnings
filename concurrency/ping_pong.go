package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func player(ctx context.Context, name string, inCh <-chan int, outCh chan<- int,
	wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case i := <-inCh:
			fmt.Println(name, ": ", i)
			time.Sleep(time.Duration(rand.Intn(10)) * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println(name, " won the game!")
				return
			case outCh <- i:
			}
		}
	}
}

func pingPong() {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	ctx2 := ctx
	defer cancel()
	var wg sync.WaitGroup

	inCh := make(chan int)
	outCh := make(chan int)

	wg.Add(2)
	go player(ctx, "Player 1", inCh, outCh, &wg)
	go player(ctx2, "Player 2", outCh, inCh, &wg)

	inCh <- 0

	wg.Wait()

	close(inCh)
	close(outCh)
}
