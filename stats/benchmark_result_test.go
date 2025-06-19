package stats

import (
	"bytes"
	"reflect"
	"testing"
)

func TestBenchmarkResult_WriteAndRead(t *testing.T) {
	var err error
	buffer := &bytes.Buffer{}
	expected := &BenchmarkResult{
		Samples:      []Duration{1, 2, 3, 4, 5, 6, 7, 8, 7, 6, 5, 4, 3, 2, 1},
		MemoryGrowth: 34,
		Allocations:  5,
	}

	CalculateFullResultStatistics(expected)
	actual := &BenchmarkResult{}
	if _, err = expected.WriteTo(buffer); err != nil {
		t.Fatal(err)
	}

	if _, err = actual.ReadFrom(buffer); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
