package benchmark

import (
	"github.com/smarty/benchy/options"
	"github.com/smarty/benchy/stats"
)

// Entry is a benchmark function, its name, and benchmark results.
type Entry struct {
	// Name is the function name.
	Name string

	// Setup is run before the benchmark.
	Setup func()

	// BenchmarkFunction is the actual function that is benchmarked.
	BenchmarkFunction func()

	// Cleanup is run after the benchmark.
	Cleanup func()

	// Overhead represents the average time it takes to run Benchy with no load.
	Overhead stats.Duration

	// Results is the actual results of the benchmark.
	Results *stats.BenchmarkResult

	// Flags describes options on this entry.
	Flags options.BenchmarkFlag
}
