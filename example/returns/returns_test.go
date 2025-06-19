package returns

import (
	"testing"

	"github.com/smarty/benchy"
	"github.com/smarty/benchy/options"
)

func BenchmarkReturns(b *testing.B) {
	var start, end int

	benchy.New(b, options.Fast).
		RegisterBenchmark("withReturnsSet", func() { start, end = withReturnsSet(10, 13) }).
		RegisterBenchmark("withReturnsUnset", func() { start, end = withReturnsUnset(10, 13) }).
		Run()

	start++
	end++
}

//go:nosplit
//go:noinline
func withReturnsSet(begin int, count int) (start int, end int) {
	start = begin
	end = begin + count
	return start, end
}

//go:nosplit
//go:noinline
func withReturnsUnset(begin int, count int) (start int, end int) {
	start = begin
	end = begin + count
	return
}
