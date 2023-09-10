package main

import (
	"fmt"
	"sync"
	"time"
)

// Implement a sliding window counter in go that has below features:
// * count different http response codes from server in last x min
// * stop the counter if the number of server error responses in last x min is more than threshold value y
// * could handle large number of responses in a short period

//   ticker       *time.Ticker
// >= 500

// what if I will create worker pool(e.g. 3) & then send the response code in another channel

type metadata struct {
	httpResponseCodeCount int
	errorCount            int
	errorThreshold        int
	lock                  sync.Mutex
}

type response struct {
	statusCode int
}

type workerPool struct {
}

func (m *metadata) readResponse(responseChan <-chan response) {
	for r := range responseChan {
		m.lock.Lock()
		if r.statusCode >= 500 {
			m.errorCount++
		} else {
			m.httpResponseCodeCount++
		}
		time.Sleep(100 * time.Millisecond)
		m.lock.Unlock()
	}
}

func (m *metadata) checkResponse(done chan bool) {
	if m.errorThreshold < m.errorCount {
		fmt.Println(m.httpResponseCodeCount, m.errorCount)
		done <- true
	}

	fmt.Println(m.httpResponseCodeCount, m.errorCount)
	m.errorCount = 0
	m.httpResponseCodeCount = 0
}

func fu_sliding_window_counter() {
	ticker := time.NewTicker(600 * time.Millisecond)
	done := make(chan bool)
	responseChan := make(chan response)
	mainDone := make(chan bool)

	m := metadata{
		errorThreshold: 5,
	}

	go m.readResponse(responseChan)

	go func() {
		for {
			select {
			case <-done:
				ticker.Stop()
				close(responseChan)
				close(done)
				mainDone <- true
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
				m.checkResponse(done)
			}
		}
	}()

	responses := []int{101, 201, 501, 202, 203, 204, 301, 501, 502, 503, 201, 300, 501,
		501, 501, 501, 501, 501, 501, 501, 501, 501}

	r := response{}

	for i := range responses {
		r.statusCode = responses[i]
		responseChan <- r
	}

	<-mainDone

	close(mainDone)
}
