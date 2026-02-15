package main

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type bidderDetails struct {
	amount int
	id     int
}

func bid(ctx context.Context, bidderID int, bidCh chan<- bidderDetails, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		bidAmount := rand.IntN(10000) + 1

		select {
		case <-ctx.Done():
			fmt.Printf("Bidder %d concluded its bidding\n", bidderID)
			return
		case bidCh <- bidderDetails{bidAmount, bidderID}:
			fmt.Printf("Bidder %d bids for amount %d\n", bidderID, bidAmount)
		}

		time.Sleep(time.Duration(rand.IntN(501)) * time.Millisecond)
	}
}

// The Fast-Paced Bidding War (Timeout vs. Early Completion)

// The Scenario: You are building an auction system. Bidders are furiously placing bids, but the auction house
// will only accept the first 5 bids or whatever bids come in before a 2-second timer expires—whichever happens first.

// The Prompt:
// Write a Go program where 3 bidder goroutines constantly attempt to send random bid amounts to an auctioneer via a shared channel.
// Each bidder waits a random amount of time (0 to 500ms) between bids.

// The auctioneer (main function) should stop the auction if either of these two conditions is met:
// It successfully collects 5 bids.
// 2 seconds have passed.
func fastPacedBidingWar() {
	runningDuration := 2 * time.Second
	noOfBiders := 3

	ch := make(chan bidderDetails)
	var wg sync.WaitGroup

	start := time.Now()
	ctx, cancel := context.WithTimeout(context.Background(), runningDuration)
	defer cancel()

	for i := range noOfBiders {
		wg.Add(1)
		go bid(ctx, i, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	totalBids := 5
	bidCount := 0

	largestBid := 0
	for bd := range ch {
		if bd.amount <= largestBid {
			continue
		}

		bidCount++
		fmt.Printf("Bidding %d placed by bidder %d\n", bd.amount, bd.id)
		if totalBids == bidCount {
			fmt.Println("Maximum limit of bidding reached... Hence closing bidding...")
			cancel()
			break
		}

		largestBid = bd.amount
	}

	fmt.Println("Auction Closed. Total Bid amount placed ", bidCount, time.Since(start))

}
