package benchmark

import (
	"sync"
	"time"
)

// RunConcurrencyBenchmark tests the CPU's ability to handle concurrent operations.
func RunConcurrencyBenchmark() float64 {
	goroutines := 100
	opsPerGoroutine := 1000000
	var wg sync.WaitGroup
	wg.Add(goroutines)

	start := time.Now()

	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < opsPerGoroutine; j++ {
				_ = j * j
			}
		}()
	}

	wg.Wait()
	duration := time.Since(start).Seconds()

	// Score is total operations divided by time.
	return float64(goroutines*opsPerGoroutine) / duration
}
