package benchmark

import (
	"time"
)

// RunMemoryBenchmark performs read/write operations to benchmark memory performance.
func RunMemoryBenchmark() float64 {
	size := 100000000
	data := make([]int, size)
	readSum := 0
	writeSum := 0

	// Write Benchmark
	start := time.Now()
	for i := range data {
		data[i] = i
		writeSum += i
	}
	writeDuration := time.Since(start).Seconds()

	// Read Benchmark
	start = time.Now()
	for _, value := range data {
		readSum += value
	}
	readDuration := time.Since(start).Seconds()

	// Use a simple metric for memory score: (read+write) operations per second
	return float64(size) / (readDuration + writeDuration)
}
