package strategies

import (
	"runtime"

	"github.com/smarty/benchy/stats"
)

// MemoryStatsStrategy calculates memory statistics.
type MemoryStatsStrategy interface {
	// SetStartingStats sets the memory usage statistics before a benchmark is
	// run.
	SetStartingStats()

	// SetEndingStats sets the memory usage statistics after a benchmark is run.
	SetEndingStats()

	// CommitStats adds the most recently calculated stats to internal values
	// for averaging and writing later.
	//
	// Parameters:
	//   - n is the number of cycles in the last benchmark run.
	CommitStats(n int)

	// WriteTo averages the values and writes the average to the
	// `result.`
	//
	// Parameters:
	//   - result is the instance to write to.
	//   - sampleCount is the number of samples taken.
	WriteTo(result *stats.BenchmarkResult, sampleCount int)
}

// ----- Active ------

type ActiveMemoryStats struct {
	memoryStats *runtime.MemStats

	allocations  float64
	memoryGrowth float64

	startAllocs uint64
	endAllocs   uint64

	startFrees uint64
	endFrees   uint64
}

func NewActiveMemoryStats() *ActiveMemoryStats {
	return &ActiveMemoryStats{
		memoryStats: &runtime.MemStats{},
	}
}

func (this *ActiveMemoryStats) SetStartingStats() {
	runtime.ReadMemStats(this.memoryStats)
	this.startAllocs = this.memoryStats.Mallocs
	this.startFrees = this.memoryStats.Frees
}

func (this *ActiveMemoryStats) SetEndingStats() {
	runtime.ReadMemStats(this.memoryStats)
	this.endAllocs = this.memoryStats.Mallocs
	this.endFrees = this.memoryStats.Frees
}

func (this *ActiveMemoryStats) CommitStats(n int) {
	this.allocations += float64(this.endAllocs-this.startAllocs) / float64(n)
	this.memoryGrowth += float64((this.endAllocs-this.endFrees)-(this.startAllocs-this.startFrees)) / float64(n)
}

func (this *ActiveMemoryStats) WriteTo(result *stats.BenchmarkResult, sampleCount int) {
	result.Allocations = this.allocations / float64(sampleCount)
	result.MemoryGrowth = this.memoryGrowth / float64(sampleCount)
}

// ----- NULL ------

type NullMemoryStats struct {
}

func NewNullMemoryStats() *NullMemoryStats {
	return &NullMemoryStats{}
}

func (this *NullMemoryStats) SetStartingStats()                                      {}
func (this *NullMemoryStats) SetEndingStats()                                        {}
func (this *NullMemoryStats) CommitStats(n int)                                      {}
func (this *NullMemoryStats) WriteTo(result *stats.BenchmarkResult, sampleCount int) {}
