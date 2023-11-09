package main

import (
	"fmt"
	"github.com/tomclarke25/cpu_benchmark/pkg/benchmark"
	"github.com/tomclarke25/cpu_benchmark/pkg/config"
    "github.com/tomclarke25/cpu_benchmark/pkg/results"
	"github.com/tomclarke25/cpu_benchmark/pkg/scoring"
)

func main() {
	fmt.Println("Starting CPU Benchmark...")

	// Load default configuration
	cfg, err := config.LoadDefaultConfig()
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
    weights := scoring.BenchmarkWeights{
        ArithmeticWeight: cfg.ArithmeticWeight.Weight,
        MemoryWeight:     cfg.MemoryWeight.Weight,
        ConcurrencyWeight: cfg.ConcurrencyWeight.Weight,
    }

	// Run benchmarks
	arithmeticScore := benchmark.RunArithmeticBenchmark()
	memoryScore := benchmark.RunMemoryBenchmark()
	concurrencyScore := benchmark.RunConcurrencyBenchmark()

	// Normalize scores
    benchmarkResults := scoring.BenchmarkResults{
        ArithmeticScore: arithmeticScore,
        MemoryScore:     memoryScore,
        ConcurrencyScore: concurrencyScore,
    }
	normalizedResults := scoring.NormalizeScores(benchmarkResults, scoring.BaselineBenchmarkResults)

	// Calculate weighted score
	totalScore := scoring.CalculateScore(normalizedResults, weights)

	// Create a map with the scores
	scores := map[string]float64{
		"Arithmetic": arithmeticScore,
		"Memory":     memoryScore,
		"Concurrency": concurrencyScore,
		"Total":      totalScore,
	}

	// Report the scores
	results.Report(scores) // This line is causing the error

	fmt.Println("CPU Benchmark completed.")

	fmt.Printf("Total Score: %f\n", totalScore)
	fmt.Println("CPU Benchmark completed.")
}