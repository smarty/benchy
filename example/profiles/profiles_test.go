package profiles

import (
	"testing"

	"github.com/smarty/benchy"
	"github.com/smarty/benchy/options"
)

func BenchmarkFast(b *testing.B) {
	benchy.New(b, options.Fast).
		RegisterBenchmark("switches", switches).
		Run()
}

func BenchmarkMedium(b *testing.B) {
	benchy.New(b, options.Medium).
		RegisterBenchmark("switches", switches).
		Run()
}

func BenchmarkFullMetrics(b *testing.B) {
	benchy.New(b, options.FullMetrics).
		RegisterBenchmark("switches", switches).
		Run()
}

func switches() {
	bools := [1_000]bool{}
	for span := 1; span <= len(bools); span++ {
		for index := span - 1; index < len(bools); index += span {
			bools[index] = !bools[index]
		}
	}
}
