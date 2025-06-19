package assertions

import (
	"testing"

	"github.com/smarty/benchy"
	"github.com/smarty/benchy/is"
	"github.com/smarty/benchy/options"
)

func BenchmarkFasterThan(b *testing.B) {
	benchy.New(b, options.Fast).
		RegisterBenchmark("fib", fib).
		RegisterBenchmark("fibWithCache", fibWithCache).
		DontPrintStats().
		Run().
		AssertThat("fibWithCache", is.FasterThan, "fib")
}
