package memory_usage

import (
	"testing"

	"github.com/smarty/benchy"
	"github.com/smarty/benchy/options"
)

var someString = ""

func BenchmarkMemoryUsage(b *testing.B) {
	benchy.New(b, options.Medium).
		RegisterBenchmark("highMemory", highMemoryUsage).
		ShowMemoryStats().
		Run()
}

func highMemoryUsage() {
	someString += "-"
	if len(someString) > 1000 {
		someString = ""
	}
}
