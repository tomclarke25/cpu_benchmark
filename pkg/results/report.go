package results

import (
	"fmt"
	"sort"
	"strings"
)

func Save(scores map[string]float64) {
	// Save the benchmark scores to a file or database.
}

func Report(scores map[string]float64) {
	// Generate a human-readable report.
	fmt.Println(strings.Repeat("=", 20))
	fmt.Println("Benchmark Results:")
	fmt.Println(strings.Repeat("=", 20))

	// Sort the keys for consistent output
	keys := make([]string, 0, len(scores))
	for k := range scores {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	var total float64
	for _, k := range keys {
		fmt.Printf("%s: %.2f\n", k, scores[k])
		total += scores[k]
	}

	avg := total / float64(len(scores))
	fmt.Printf("\nAverage Score: %.2f\n", avg)
	fmt.Println(strings.Repeat("=", 20))
}