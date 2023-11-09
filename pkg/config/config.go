package config

import (
	"encoding/json"
	"os"
)

type BenchmarkConfig struct {
	ArithmeticWeight struct {
		Weight float64 `json:"weight"`
	} `json:"arithmetic"`
	MemoryWeight struct {
		Weight float64 `json:"weight"`
	} `json:"memory"`
	ConcurrencyWeight struct {
		Weight float64 `json:"weight"`
	} `json:"concurrency"`
}

func LoadDefaultConfig() (BenchmarkConfig, error) {
	file, err := os.Open("configs/defaults.json")
	if err != nil {
		return BenchmarkConfig{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := BenchmarkConfig{}
	err = decoder.Decode(&config)
	if err != nil {
		return BenchmarkConfig{}, err
	}

	return config, nil
}
