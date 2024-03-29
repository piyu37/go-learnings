package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

// This benchmark tells the performance of worker pool when we use one go-routine or multiple
// BenchmarkWithNoBuffer-12: number 12 represents the number of goroutines being used during the benchmark
// 1 col: no. of iterations
// 2 col: average time per iteration
// 3 col: average number of bytes allocated per iteration. It indicates the memory allocation size of each iteration.
// 4 col: The average number of memory allocations per iteration
func BenchmarkWithNoBuffer(b *testing.B) {
	benchmarkWithBuffer(b, 0)
}

func BenchmarkWithBufferSizeOf1(b *testing.B) {
	benchmarkWithBuffer(b, 1)
}

func BenchmarkWithBufferSizeEqualsToNumberOfWorker(b *testing.B) {
	benchmarkWithBuffer(b, 5)
}

func BenchmarkWithBufferSizeExceedsNumberOfWorker(b *testing.B) {
	benchmarkWithBuffer(b, 25)
}

func benchmarkWithBuffer(b *testing.B, size int) {
	for i := 0; i < b.N; i++ {
		c := make(chan uint32, size)

		var wg sync.WaitGroup
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := uint32(0); i < 1000; i++ {
				c <- i % 2
			}
			close(c)
		}()

		var total uint32
		for w := 0; w < 5; w++ {
			wg.Add(1)
			go func() {
				defer wg.Done()

				for {
					v, ok := <-c
					if !ok {
						break
					}
					atomic.AddUint32(&total, v)
				}
			}()
		}

		wg.Wait()
	}
}
