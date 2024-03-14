package main

import (
	"fmt"
	"sync"
	"time"
)

type metadata struct {
	httpResponseCodeCount int
	errorCount            int
	errorThreshold        int
	lock                  sync.Mutex
}

type response struct {
	statusCode int
}

func (m *metadata) readResponse(responseChan <-chan response, wg *sync.WaitGroup) {
	defer wg.Done()

	for r := range responseChan {
		m.lock.Lock()
		fmt.Println(time.Now())
		if r.statusCode >= 500 {
			m.errorCount++
		} else {
			m.httpResponseCodeCount++
		}
		m.lock.Unlock()
		time.Sleep(250 * time.Millisecond)
	}
}

func (m *metadata) checkResponse() bool {
	m.lock.Lock()
	if m.errorThreshold < m.errorCount {
		fmt.Println("response: ", m.httpResponseCodeCount, m.errorCount)
		m.lock.Unlock()
		return true
	}

	fmt.Println("response: ", m.httpResponseCodeCount, m.errorCount)
	m.errorCount = 0
	m.httpResponseCodeCount = 0
	m.lock.Unlock()

	return false
}

func (m *metadata) produceEvents(timer *time.Timer, responses []int, ch chan<- response,
	cronDone, prodDone chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	r := response{}

	for i := 0; i < len(responses); i++ {
		select {
		case <-timer.C:
			cronDone <- true
			fmt.Println("closing: time deadline crossed")
			return
		case <-prodDone:
			return
		default:
			r.statusCode = responses[i]
			ch <- r
		}
	}

	cronDone <- true
	fmt.Println("closing: all data read")
}

func (m *metadata) cronJob(tickDuration int, timer *time.Timer, cronDone, prodDone chan bool,
	wg *sync.WaitGroup) {
	defer wg.Done()
	ticker := time.NewTicker(time.Duration(tickDuration) * time.Millisecond)
	for {
		select {
		case t := <-ticker.C:
			fmt.Println("Tick at", t)
			if m.checkResponse() {
				prodDone <- true
				fmt.Println("closing: error threshold crossed")
				return
			}
		case <-cronDone:
			return
		}
	}
}

// Implement a sliding window counter in go that has below features:
// * count different http response codes from server in last x min
// * stop the counter if the number of server error responses in last x min is more than threshold value y
// * could handle large number of responses in a short period

//   ticker       *time.Ticker
// >= 500

// what if I will create worker pool(e.g. 3) & then send the response code in another channel
// 1. go routine which takes care reading values and sending it;
// use array to read/randomly generate status code.
// This will stop once we will get more error than threshold when checking at some
// interval. Or close after some y time. (1 ch multiple producers and consumers)
// 2. checkResponse
// 3. ReadResponse
func fu_sliding_window_counter() {
	start := time.Now()
	y := 5               // in second
	tickDuration := 1000 // in millisecond
	timer := time.NewTimer(time.Duration(y) * time.Second)
	responseChan := make(chan response, 2) // buffer channel
	var wg sync.WaitGroup
	cronDone := make(chan bool)
	prodDone := make(chan bool)
	consumerWorkeres := 3

	m := metadata{
		errorThreshold: 7,
	}

	responses := []int{101, 201, 501, 202, 203, 204, 301, 501, 502, 503, 201, 300, 501,
		501, 501, 501, 501, 501, 501, 501, 501, 501, 501} // total 23

	wg.Add(1)
	go m.produceEvents(timer, responses, responseChan, cronDone, prodDone, &wg)
	for i := 0; i < consumerWorkeres; i++ {
		wg.Add(1)
		go m.readResponse(responseChan, &wg)
	}

	wg.Add(1)
	go m.cronJob(tickDuration, timer, cronDone, prodDone, &wg)

	wg.Wait()
	m.checkResponse()

	close(cronDone)
	close(prodDone)

	fmt.Println("time: ", time.Since(start))
}
