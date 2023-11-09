package benchmark

import (
	"math"
	"math/rand"
	"sort"
	"time"
)

// RunArithmeticBenchmark performs intense arithmetic operations to benchmark CPU performance.
func RunBasicArithmeticBenchmark() float64 {
	start := time.Now()
	ops := 0

	for i := 0; i < 100000000; i++ {
		_ = math.Sqrt(float64(i))
		ops++
	}

	duration := time.Since(start).Seconds()
	return float64(ops) / duration // operations per second
}


// RunDetailedArithmeticBenchmark performs a variety of arithmetic operations to benchmark CPU performance.
func RunDetailedArithmeticBenchmark() float64 {
    start := time.Now()
    ops := 0

    for i := 1; i <= 1000000000; i++ { // Increase the loop count
        // Perform a variety of arithmetic operations
        _ = i + i
        _ = i - 1
        _ = i * i
        _ = i / 2
        _ = math.Sqrt(float64(i))

        ops += 6 // increment the operations count by 6 for each loop iteration
    }

    duration := time.Since(start).Seconds()
    return float64(ops) / duration // operations per second
}

// RunComplexBenchmark performs a complex operation (sorting a large array) to benchmark CPU performance.
func RunComplexBenchmark() float64 {
    start := time.Now()

    // Generate a large array of random numbers
    nums := make([]int, 1000000000)
    for i := range nums {
        nums[i] = rand.Int()
    }

	// Sort the array
	sort.Ints(nums)

	duration := time.Since(start).Seconds()

	// The benchmark result is the time taken to sort the array
	return duration
}

// RunAllBenchmarks runs all benchmarks and returns the average score.
func RunArithmeticBenchmark() float64 {
	arithmeticBenchmark := RunBasicArithmeticBenchmark()
	detailedArithmeticBenchmark := RunDetailedArithmeticBenchmark()
	complexBenchmark := RunComplexBenchmark()

	// Calculate the average
	average := (arithmeticBenchmark + detailedArithmeticBenchmark + complexBenchmark) / 3

	return average
}
