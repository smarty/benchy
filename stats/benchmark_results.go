package stats

import (
	"encoding/binary"
	"io"
	"strings"
	"testing"
)

// BenchmarkResults is a collection of BenchmarkResult that can be inspected
// together.
type BenchmarkResults struct {
	tb testing.TB

	// Collection is the direct accessor for the collection of BenchmarkResult.
	Collection []*BenchmarkResult
}

// NewBenchmarkResults generates a new collection of results that can be
// inspected together.
func NewBenchmarkResults(tb testing.TB) *BenchmarkResults {
	return &BenchmarkResults{
		tb: tb,
	}
}

// WriteTo fulfills the io.WriterTo interface.
func (this *BenchmarkResults) WriteTo(writer io.Writer) (count int64, err error) {
	var (
		n        int
		subCount int64
	)

	collectionSizeBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(collectionSizeBuffer, uint32(len(this.Collection)))
	n, err = writer.Write(collectionSizeBuffer)
	count += int64(n)
	if err != nil {
		return count, err
	}

	for _, result := range this.Collection {
		subCount, err = result.WriteTo(writer)
		count += subCount
		if err != nil {
			return count, err
		}
	}

	return count, nil
}

// ReadFrom fulfills the io.ReaderFrom interface.
func (this *BenchmarkResults) ReadFrom(reader io.Reader) (count int64, err error) {
	var (
		n        int
		subCount int64
	)

	collectionSizeBuffer := make([]byte, 4)
	n, err = reader.Read(collectionSizeBuffer)
	count += int64(n)
	if err != nil {
		return count, err
	}

	collectionSize := binary.LittleEndian.Uint32(collectionSizeBuffer)

	for i := 0; i < int(collectionSize); i++ {
		result := &BenchmarkResult{}
		subCount, err = result.ReadFrom(reader)
		count += subCount
		if err != nil {
			return count, err
		}

		this.Collection = append(this.Collection, result)
	}

	return count, nil
}

// AssertThat tests a specified condition on one or more benchmarks.
func (this *BenchmarkResults) AssertThat(left string, operator TestOperator, right ...string) *BenchmarkResults {
	var (
		leftResult   *BenchmarkResult
		rightResults = make([]*BenchmarkResult, len(right))
	)

	for _, result := range this.Collection {
		if strings.EqualFold(left, result.Name) {
			leftResult = result
		}

		for iRightName, rightName := range right {
			if strings.EqualFold(rightName, result.Name) {
				rightResults[iRightName] = result
			}
		}
	}

	exitEarly := false
	if leftResult == nil {
		this.tb.Errorf("assertion error: %v", generateBenchmarkNotFoundError(left))
		exitEarly = true
	}

	for iResult, rightResult := range rightResults {
		if rightResult == nil {
			this.tb.Errorf("assertion error: %v", generateBenchmarkNotFoundError(right[iResult]))
			exitEarly = true
		}
	}

	if exitEarly {
		return this
	}

	err := operator(leftResult, rightResults...)
	if err != nil {
		this.tb.Errorf("assertion error: %v", err.Error())
	}

	return this
}
