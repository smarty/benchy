package options

// BenchmarkProfile defines some defaults for a benchmark.
type BenchmarkProfile int

const (
	// Fast specifies that the benchmark(s) don't need many samples, no readout
	// required. Think unit test, though not limited to unit test like
	// benchmarks exclusively.
	//
	// Defaults are: 3 samples, Long = off.
	Fast BenchmarkProfile = iota

	// Medium specifies that the benchmark(s) could be either sort or long,
	// a low resolution statistics report is desired. Think integration test,
	// though not limited to integration like benchmarks exclusively.
	//
	// Defaults are: 10 samples, Long = off.
	Medium

	// FullMetrics specifies that the benchmark(s) will be long-running and a
	// full readout with high resolution is desired. Think end-to-end test,
	// though not limited to end-to-end like benchmarks exclusively.
	//
	// Defaults are: 25 samples, Long = on.
	FullMetrics
)
