package benchmark

import (
	"fmt"
	"testing"

	"github.com/smarty/benchy/internal/benchmark/strategies"
	"github.com/smarty/benchy/options"
	"github.com/smarty/benchy/stats"
)

// SampleOverhead runs Benchy with no benchmark function in order to record the
// average overhead.
//
// Parameters:
//   - b is the benchmark runner.
//   - entry contains benchmark information.
//   - sampleCount is the number of samples for this benchmark.
func SampleOverhead(b *testing.B, entry *Entry, sampleCount int) {
	if !entry.Flags.Contains(options.OverheadSampling) {
		entry.Overhead = 0
		return
	}

	overHeadEntry := &Entry{
		Setup:             entry.Setup,
		BenchmarkFunction: func() {},
		Cleanup:           entry.Cleanup,
	}
	result := sampleHelper(b, fmt.Sprintf("%s [overhead]", entry.Name), overHeadEntry, min(sampleCount, 5), stats.Duration(0))
	stats.CalculateAverage(result)
	entry.Overhead = result.Average
}

// Sample runs the benchmark function `sampleCount` times and records the sample
// durations.
//
// Parameters:
//   - b is the benchmark runner.
//   - entry contains benchmark information.
//   - sampleCount is the number of samples for this benchmark.
func Sample(b *testing.B, entry *Entry, sampleCount int) {
	result := sampleHelper(b, entry.Name, entry, sampleCount, entry.Overhead)
	stats.CalculateFullResultStatistics(result)

	entry.Results = result
}

func sampleHelper(b *testing.B, name string, entry *Entry, sampleCount int, overhead stats.Duration) *stats.BenchmarkResult {
	var (
		previousN   int
		finalSample stats.Duration

		memoryStats strategies.MemoryStatsStrategy
		pprofCPU    strategies.PProfCPUStrategy = strategies.NewNullPProfCPU()
	)

	result := &stats.BenchmarkResult{
		Name:     name,
		Samples:  make([]stats.Duration, 0, sampleCount),
		Outliers: make([]stats.Duration, 0, sampleCount/2),
	}

	memoryStats = strategies.NewActiveMemoryStats()
	if entry.Flags.Contains(options.PProfCPU) {
		pprofCPU = strategies.NewActivePProfCPU(b, name)
	}

	for sample := 0; sample < sampleCount; sample++ {
		finalSample = 0
		previousN = 0

		b.Run(name, func(b *testing.B) {
			pprofCPU.StartRecording()
			memoryStats.SetStartingStats()
			entry.Setup()
			for i := 0; i < b.N; i++ {
				entry.BenchmarkFunction()
			}

			if b.N < previousN {
				result.Samples = append(result.Samples, finalSample)
			}

			previousN = b.N
			finalSample = stats.Duration(b.Elapsed().Nanoseconds()) / stats.Duration(b.N)
			memoryStats.SetEndingStats()
			pprofCPU.StopRecording()
			entry.Cleanup()
		})

		pprofCPU.WriteRecording()
		result.Samples = append(result.Samples, max(0, finalSample-overhead))
		memoryStats.CommitStats(previousN)
	}

	pprofCPU.WriteRunnerScript()
	memoryStats.WriteTo(result, sampleCount)
	return result
}
