package with_data

import (
	"testing"

	"github.com/smarty/benchy"
	"github.com/smarty/benchy/options"
	"github.com/smarty/benchy/providers"
)

func BenchmarkWithData(b *testing.B) {
	fibBuilder_5And10 := providers.New1(fib)
	fibBuilder_5And10.Add(5)
	fibBuilder_5And10.Add(10)
	fib_5And10 := fibBuilder_5And10.BenchmarkFunc()

	fibBuilder_25And30 := providers.New1(fib)
	fibBuilder_25And30.Add(25)
	fibBuilder_25And30.Add(30)
	fib_25And30 := fibBuilder_25And30.BenchmarkFunc()

	benchy.New(b, options.Fast).
		RegisterBenchmark("fib5And10", fib_5And10).
		RegisterBenchmark("fib25And30", fib_25And30).
		Run()
}

func fib(n int) {
	_fib(n)
}

func _fib(n int) int {
	if n <= 1 {
		return 1
	}

	return _fib(n-1) + _fib(n-2)
}
