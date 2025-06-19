package benchy

import (
	"os"
	"strings"
	"testing"

	"github.com/smarty/benchy/internal/benchmark"
	"github.com/smarty/benchy/internal/params"
	"github.com/smarty/benchy/internal/rendering"
	"github.com/smarty/benchy/options"
	"github.com/smarty/benchy/stats"
)

// Benchy is a robust benchmarking service.
type Benchy struct {
	b               *testing.B
	benchmarks      []*benchmark.Entry
	printer         statPrinter
	printMemoryFunc rendering.ExtraRenderingFunc
	profile         options.BenchmarkProfile
	sampleCount     int
	runningLong     bool
}

// New sets up a new Benchy, ready to benchmark some code.
//
// Parameters:
//   - b is the required *testing.B for running the benchmarks.
//   - profile tells Benchy what defaults to use. See [options.BenchmarkProfile]
//     for all options and in-depth descriptions of each. Generally speaking,
//     [options.Fast] is best for a quick benchmark, whereas [options.Medium] is
//     best for comparing two strategies.
func New(b *testing.B, profile options.BenchmarkProfile) *Benchy {
	return &Benchy{
		b:           b,
		printer:     new(activePrinter),
		runningLong: !testing.Short(),
		profile:     profile,
	}
}

// DontPrintStats turns off printing stats. This is useful for automated systems
// which are not going to examine the printout. Otherwise, you generally have
// no reason ot turn off stats.
func (this *Benchy) DontPrintStats() *Benchy {
	this.printer = new(nullPrinter)
	return this
}

// SetSampleCount sets the number of samples that will be taken. Default is
// controlled by the profile chosen when calling [benchy.New].
//
// Parameters:
//   - sampleCount is the number of samples that Benchy will run. The minimum
//     sample count is 1. If the flag `-test.samples` is set from the CLI, then
//     that flag definition will take precedence over this method.
func (this *Benchy) SetSampleCount(sampleCount int) *Benchy {
	this.sampleCount = sampleCount
	return this
}

// ShowMemoryStats activates the rendering of memory statistics.
//
// Benchy must have a sample count of at least stats.MinFullCalculation to show
// any statistics.
func (this *Benchy) ShowMemoryStats() *Benchy {
	this.printMemoryFunc = rendering.RenderMemoryFunc
	return this
}

// RegisterBenchmark adds a new function to be benchmarked.
//
// Parameters:
//   - name is a unique identifier for the registered function. If the name is
//     not unique among registered functions, then the benchmark will fail.
//   - benchmarkFunction must fulfil the niladic definition `func()` with no
//     returns. Review the examples below for specific cases.
//   - flags sets any number of options for this benchmark function. See
//     [options.BenchmarkFlag] for details on what options are available and
//     better descriptions on what they do.
//
// Examples:
//
// A function with return values, wrap the function:
//
//	benchy.New(b, options.Medium).
//	RegisterBenchmark("with returns", func() {myFunc()}).
//	Run()
//
// A function with parameters that can always be the same, wrap the function:
//
//	benchy.New(b, options.Medium).
//	RegisterBenchmark("with parameters", func() {myFunc(arg1, arg2, ...)}).
//	Run()
//
// A function with parameters that should have different values each sample run,
// use a data provider:
//
//	dp := providers.NewX(myFunc). // X is the number of input parameters between 1 and 6.
//	Add(arg1_1, arg1_2, ...).
//	Add(arg2_1, arg2_2, ...).
//	Add(arg3_1, arg3_2, ...)
//	benchy.New(b, options.Medium).
//	RegisterBenchmark("dynamic parameters", dp.WrapBenchmarkFunc(myFunc)).
//	Run()
func (this *Benchy) RegisterBenchmark(name string, benchmarkFunction func(), flags ...options.BenchmarkFlag) *Benchy {
	for _, entry := range this.benchmarks {
		if strings.EqualFold(entry.Name, name) {
			this.b.Errorf(
				"registering a second benchmark with similar name: previously registered '%s', now registering '%s'",
				entry.Name,
				name)
		}
	}

	benchmarkEntry := &benchmark.Entry{
		Name:              name,
		Setup:             func() {},
		BenchmarkFunction: benchmarkFunction,
		Cleanup:           func() {},
		Results:           new(stats.BenchmarkResult),
	}

	benchmarkEntry.Flags.Set(flags...)
	if this.profile == options.FullMetrics {
		benchmarkEntry.Flags.Set(options.Long)
	}

	this.benchmarks = append(this.benchmarks, benchmarkEntry)
	return this
}

// RegisterSetup adds a setup function to the already named and registered
// benchmark. Setup functions will run on Benchy sample (See [SetSampleCount]).
// When a setup is registered for a function, it will automatically turn on
// overhead sampling, which subtracts the runtime of the setup function from
// the total runtime to get a more accurate value.
//
// Parameters:
//   - benchmarkName must be the identifier for a benchmark function that has
//     already been registered.
//   - setupFunction will be run once every sampling before the proper benchmark
//     function runs.
func (this *Benchy) RegisterSetup(benchmarkName string, setupFunction func()) *Benchy {
	registered := false
	for _, entry := range this.benchmarks {
		if !strings.EqualFold(entry.Name, benchmarkName) {
			continue
		}

		entry.Setup = setupFunction
		entry.Flags.Set(options.OverheadSampling)
		registered = true
		break
	}

	if !registered {
		this.b.Errorf(
			"registering a setup function for '%s' failed, this benchmark has not yet been registered",
			benchmarkName)
	}

	return this
}

// RegisterCleanup adds a cleanup function to the already named and registered
// benchmark. Cleanup functions will run on Benchy sample (See
// [SetSampleCount]). When a cleanup is registered for a function, it will
// automatically turn on overhead sampling, which subtracts the runtime of the
// cleanup function from the total runtime to get a more accurate value.
//
// Parameters:
//   - benchmarkName must be the identifier for a benchmark function that has
//     already been registered.
//   - cleanupFunction will be run once every sampling after the proper
//     benchmark function runs.
func (this *Benchy) RegisterCleanup(benchmarkName string, cleanupFunction func()) *Benchy {
	registered := false
	for _, entry := range this.benchmarks {
		if !strings.EqualFold(entry.Name, benchmarkName) {
			continue
		}

		entry.Cleanup = cleanupFunction
		entry.Flags.Set(options.OverheadSampling)
		registered = true
		break
	}

	if !registered {
		this.b.Errorf(
			"registering a cleanup function for '%s' failed, this benchmark has not yet been registered",
			benchmarkName)
	}

	return this
}

// Run runs all the registered benchmarks and returns the results.
//
// Returns:
//   - benchmarkResults is a collection of all the results of the various
//     benchmarks that have run. See [stats.BenchmarkResults] for more details.
func (this *Benchy) Run() (benchmarkResults *stats.BenchmarkResults) {
	this.sampleCount = params.SelectSampleCount(this.sampleCount, this.profile, os.Args)
	for _, entry := range this.benchmarks {
		if entry.Flags.Contains(options.Long) && !this.runningLong {
			continue
		}

		benchmark.SampleOverhead(this.b, entry, this.sampleCount)
		benchmark.Sample(this.b, entry, this.sampleCount)
	}

	results := make([]*stats.BenchmarkResult, 0, len(this.benchmarks))
	for _, entry := range this.benchmarks {
		if entry.Flags.Contains(options.Long) && !this.runningLong {
			continue
		}

		this.printer.printHistogram(entry.Results, this.sampleCount)
		results = append(results, entry.Results)
	}

	if len(results) > 0 {
		this.printer.printReportCard(results, this.sampleCount, this.printMemoryFunc)
	}

	benchmarkResults = stats.NewBenchmarkResults(this.b)
	benchmarkResults.Collection = results
	return benchmarkResults
}
