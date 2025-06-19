package stats

import (
	"encoding/binary"
	"io"
	"math"

	"github.com/smarty/benchy/internal/statistics"
)

// The MinFullCalculation sets a sample threshold for full statistics
// calculation. If the sample count is equal to or lower than this number,
// there is not enough data to run full statistics calculations.
const MinFullCalculation = 10

// TestOperator defines the assert operator.
type TestOperator func(left *BenchmarkResult, right ...*BenchmarkResult) error

// BenchmarkResult contains performance metrics useful for comparisons.
type BenchmarkResult struct {
	// Samples is a collection of all aggregated benchmark sample timings,
	// including Outliers.
	Samples []Duration

	// Outliers is a collection of all the outliers from Samples.
	Outliers []Duration

	// Name is the provided name for the benchmark.
	Name string

	// Average represents the simple median of the Samples excluding Outliers.
	Average Duration

	// Median is the "middle" sample, excluding Outliers.
	Median Duration

	// Max is the longest single sample, excluding Outliers.
	Max Duration

	// Min is the shortest single sample, excluding Outliers.
	Min Duration

	// StandardDeviation is the standard deviation from Samples excluding
	// Outliers.
	StandardDeviation Duration

	// StandardError is the standard error from Samples excluding Outliers.
	StandardError Duration

	// Histogram is a collection of result-statistics organized in buckets to
	// measure and graph modality.
	Histogram []int

	// Modality is the number of peaks in the histogram data.
	Modality int

	// FourSigma is the 99.99% upper bound.
	// When modality is greater than 0, this value is based on the highest mode.
	// When modality is equal to 1, the median is used.
	FourSigma Duration

	// Allocations is the average number of allocations per operation.
	Allocations float64

	// MemoryGrowth is the average number of allocations per operation that are
	// not freed.
	MemoryGrowth float64
}

// WriteTo fulfills the io.WriterTo interface.
func (this *BenchmarkResult) WriteTo(writer io.Writer) (count int64, err error) {
	var (
		n int
	)

	nameLength := int32(len(this.Name))
	nameLengthBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(nameLengthBuffer, uint32(nameLength))
	n, err = writer.Write(nameLengthBuffer)
	count += int64(n)
	if err != nil {
		return count, err
	}

	n, err = writer.Write([]byte(this.Name))
	count += int64(n)
	if err != nil {
		return count, err
	}

	sampleCount := int32(len(this.Samples))
	sampleCountBuffer := make([]byte, 4)
	binary.LittleEndian.PutUint32(sampleCountBuffer, uint32(sampleCount))
	n, err = writer.Write(sampleCountBuffer)
	count += int64(n)
	if err != nil {
		return count, err
	}

	for _, sample := range this.Samples {
		bits := math.Float64bits(float64(sample))
		count += 8
		err = binary.Write(writer, binary.LittleEndian, bits)
		if err != nil {
			return count, err
		}
	}

	memoryGrowthBits := math.Float64bits(this.MemoryGrowth)
	memoryGrowthBuffer := make([]byte, 8)
	binary.LittleEndian.PutUint64(memoryGrowthBuffer, memoryGrowthBits)
	n, err = writer.Write(memoryGrowthBuffer)
	count += int64(n)
	if err != nil {
		return count, err
	}

	allocationsBits := math.Float64bits(this.Allocations)
	allocationsBuffer := make([]byte, 8)
	binary.LittleEndian.PutUint64(allocationsBuffer, allocationsBits)
	n, err = writer.Write(allocationsBuffer)
	count += int64(n)
	if err != nil {
		return count, err
	}

	return count, nil
}

// ReadFrom fulfills the io.ReaderFrom interface.
func (this *BenchmarkResult) ReadFrom(reader io.Reader) (count int64, err error) {
	var (
		n int
	)

	nameLengthBuffer := make([]byte, 4)
	n, err = reader.Read(nameLengthBuffer)
	count += int64(n)
	if err != nil {
		return count, err
	}

	nameLength := binary.LittleEndian.Uint32(nameLengthBuffer)
	nameBuffer := make([]byte, nameLength)
	n, err = reader.Read(nameBuffer)
	count += int64(n)
	if err != nil {
		return count, err
	}

	this.Name = string(nameBuffer)

	sampleCountBuffer := make([]byte, 4)
	n, err = reader.Read(sampleCountBuffer)
	count += int64(n)
	if err != nil {
		return count, err
	}

	sampleCount := binary.LittleEndian.Uint32(sampleCountBuffer)

	for i := 0; i < int(sampleCount); i++ {
		bits := uint64(0)
		count += 8
		err = binary.Read(reader, binary.LittleEndian, &bits)
		if err != nil {
			return count, err
		}

		sample := math.Float64frombits(bits)
		this.Samples = append(this.Samples, Duration(sample))
	}

	memoryGrowthBuffer := make([]byte, 8)
	n, err = reader.Read(memoryGrowthBuffer)
	count += int64(n)
	if err != nil {
		return count, err
	}

	memoryGrowthBits := binary.LittleEndian.Uint64(memoryGrowthBuffer)
	this.MemoryGrowth = math.Float64frombits(memoryGrowthBits)

	allocationsBuffer := make([]byte, 8)
	n, err = reader.Read(allocationsBuffer)
	count += int64(n)
	if err != nil {
		return count, err
	}

	allocationsBits := binary.LittleEndian.Uint64(allocationsBuffer)
	this.Allocations = math.Float64frombits(allocationsBits)

	CalculateFullResultStatistics(this)
	return count, nil
}

// CalculateFullResultStatistics calculates all the statistics for the result.
func CalculateFullResultStatistics(result *BenchmarkResult) {
	if len(result.Samples) < MinFullCalculation {
		CalculateAverage(result)
	}

	coreSamples, outliers := statistics.SeparateOutliers(result.Samples)
	result.Outliers = outliers
	result.Min, result.Max = statistics.MinMax(result.Samples)
	result.Average = statistics.Average(coreSamples)
	result.Median = statistics.Median(coreSamples)
	result.StandardError = statistics.StandardError(coreSamples)
	result.StandardDeviation = statistics.StandardDeviation(coreSamples)
	result.Histogram = statistics.Histogram(result.Samples)
	result.Modality = statistics.ModalityFromHistogram(result.Histogram)
	result.FourSigma = statistics.FourSigma(result.Samples, result.StandardDeviation)
}

// CalculateAverage only calculates the average statistic for this result.
func CalculateAverage(result *BenchmarkResult) {
	result.Average = statistics.Average(result.Samples)
}
