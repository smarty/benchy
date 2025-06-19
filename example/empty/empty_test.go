package empty

import (
	"testing"

	"github.com/smarty/benchy"
	"github.com/smarty/benchy/options"
)

func BenchmarkEmpty(b *testing.B) {
	benchy.New(b, options.Fast).
		RegisterBenchmark("nothing", func() {}, options.OverheadSampling).
		Run()
}
