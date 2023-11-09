// pkg/scoring/scoring.go

package scoring

// BenchmarkResults holds the raw results from various benchmarks.
type BenchmarkResults struct {
	ArithmeticScore   float64
	MemoryScore       float64
	ConcurrencyScore  float64
}

// BenchmarkWeights holds the weights for each benchmark result.
// These could also be configurable via a config file.
type BenchmarkWeights struct {
	ArithmeticWeight   float64
	MemoryWeight       float64
	ConcurrencyWeight  float64
}

// DefaultBenchmarkWeights provides a default set of weights.
// In a real system, you might allow these to be adjusted by user input or configuration files.
var DefaultBenchmarkWeights = BenchmarkWeights{
	ArithmeticWeight:   0.25,
	MemoryWeight:       0.25,
	ConcurrencyWeight:  0.25,
}

// CalculateScore computes the weighted score based on the benchmark results and weights.
func CalculateScore(results BenchmarkResults, weights BenchmarkWeights) float64 {
	totalScore := results.ArithmeticScore*weights.ArithmeticWeight +
		results.MemoryScore*weights.MemoryWeight +
		results.ConcurrencyScore*weights.ConcurrencyWeight

	return totalScore
}

// BaselineBenchmarkResults defines the scores of the baseline system.
var BaselineBenchmarkResults = BenchmarkResults{
	ArithmeticScore:   1000.0, // Baseline arithmetic score
	MemoryScore:       1000.0, // Baseline memory score
	ConcurrencyScore:  1000.0, // Baseline concurrency score
}

// NormalizeScores adjusts the raw scores based on the baseline results.
func NormalizeScores(results BenchmarkResults, baseline BenchmarkResults) BenchmarkResults {
	return BenchmarkResults{
		ArithmeticScore:   results.ArithmeticScore / baseline.ArithmeticScore,
		MemoryScore:       results.MemoryScore / baseline.MemoryScore,
		ConcurrencyScore:  results.ConcurrencyScore / baseline.ConcurrencyScore,
	}
}

