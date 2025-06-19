package memory_leak

import (
	"testing"

	"github.com/smarty/benchy"
	"github.com/smarty/benchy/options"
)

var (
	head    *chainLink
	current *chainLink
)

type chainLink struct {
	next *chainLink
}

func BenchmarkMemoryLeak(b *testing.B) {
	benchy.New(b, options.Medium).
		SetSampleCount(10).
		RegisterBenchmark("memoryLeak", memoryLeak, options.Long).
		RegisterSetup("memoryLeak", func() {
			head = &chainLink{}
			current = head
		}).
		ShowMemoryStats().
		Run()
}

func memoryLeak() {
	current.next = &chainLink{}
	current = current.next
}
