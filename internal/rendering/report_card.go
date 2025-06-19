package rendering

import (
	"fmt"
	"strings"

	"github.com/smarty/benchy/stats"
)

// ExtraRenderingFunc is a function that is used for rendering extra columns.
type ExtraRenderingFunc func(*[][]string, []*stats.BenchmarkResult)

// ReportCard renders the report-card as a series of lines which can be written out.
//
// Ansi codes are used to color the text.
func ReportCard(results []*stats.BenchmarkResult, sampleCount int, memFunc ExtraRenderingFunc) []string {
	data := make([][]string, 0)

	addBenchmarkNames(&data, results)
	addColumn(&data, "AVERAGE", results, func(result *stats.BenchmarkResult) stats.Duration { return result.Average })
	if sampleCount >= stats.MinFullCalculation {
		addColumn(&data, "MEDIAN", results, func(result *stats.BenchmarkResult) stats.Duration { return result.Median })
		addColumn(&data, "MIN", results, func(result *stats.BenchmarkResult) stats.Duration { return result.Min })
		addColumn(&data, "MAX", results, func(result *stats.BenchmarkResult) stats.Duration { return result.Max })
		addColumn(&data, "STD DEV", results, func(result *stats.BenchmarkResult) stats.Duration { return result.StandardDeviation })
		addColumn(&data, "STD ERR", results, func(result *stats.BenchmarkResult) stats.Duration { return result.StandardError })
		addColumn(&data, "4σ", results, func(result *stats.BenchmarkResult) stats.Duration { return result.FourSigma })
		if memFunc != nil {
			memFunc(&data, results)
		}
	}

	// add two lines for the table header
	lines := make([]string, len(results)+2)
	sb := strings.Builder{}
	for iLine := range lines {
		sb.Reset()
		sb.WriteString(ansi_blue)
		for iColumn, column := range data {
			if iColumn > 0 {
				// header separator
				if iLine == 1 {
					sb.WriteString("-+-")
				} else {
					sb.WriteString(" | ")
				}
			}

			sb.WriteString(column[iLine])
		}

		sb.WriteString(ansi_reset)
		lines[iLine] = sb.String()
	}

	return lines
}

// RenderMemoryFunc satisfies the ExtraRenderingFunc interface for rendering memory statistics.
func RenderMemoryFunc(data *[][]string, results []*stats.BenchmarkResult) {
	addColumnFloat(data, "ALLOCATIONS", results, func(result *stats.BenchmarkResult) float64 { return result.Allocations })
	addColumnFloat(data, "MEMORY GROWTH", results, func(result *stats.BenchmarkResult) float64 { return result.MemoryGrowth })
}

func addBenchmarkNames(data *[][]string, results []*stats.BenchmarkResult) {
	nameLength := stringLength("BENCHMARK")
	for _, result := range results {
		nameLength = max(nameLength, stringLength(result.Name))
	}

	// add two lines for the table header
	column := make([]string, len(results)+2)
	column[0] = padRight("BENCHMARK", nameLength, ' ')
	column[1] = padRight("", nameLength, '-')
	for iResult, result := range results {
		column[iResult+2] = padRight(result.Name, nameLength, ' ')
	}

	*data = append(*data, column)
}

func addColumn(data *[][]string, columnName string, results []*stats.BenchmarkResult, getField func(result *stats.BenchmarkResult) stats.Duration) {
	// add two lines for the table header
	column := make([]string, len(results)+2)
	length, unit := calculateRecommendedReportItemLength(columnName, results, getField)
	column[0] = padLeft(columnName, length, ' ')
	column[1] = padLeft("", length, '-')
	for iResult, result := range results {
		value := getField(result)
		column[iResult+2] = padLeft(value.RenderWithUnit(unit), length, ' ')
	}

	*data = append(*data, column)
}

func addColumnFloat(data *[][]string, columnName string, results []*stats.BenchmarkResult, getField func(result *stats.BenchmarkResult) float64) {
	// add two lines for the table header
	column := make([]string, len(results)+2)
	length := calculateRecommendedReportItemLengthFloat(columnName, results, getField)
	column[0] = padLeft(columnName, length, ' ')
	column[1] = padLeft("", length, '-')
	for iResult, result := range results {
		value := getField(result)
		column[iResult+2] = padLeft(fmt.Sprintf("%0.3f", value), length, ' ')
	}

	*data = append(*data, column)
}

func calculateRecommendedReportItemLength(columnName string, results []*stats.BenchmarkResult, getField func(result *stats.BenchmarkResult) stats.Duration) (length int, unit string) {
	units := make([]string, 0, len(results))
	for _, result := range results {
		duration := getField(result)
		units = append(units, duration.Unit())
	}

	smallestUnit := stats.SmallestUnit(units...)

	length = stringLength(columnName)
	for _, result := range results {
		duration := getField(result)
		length = max(length, duration.RenderLengthAsUnit(smallestUnit))
	}

	return length, smallestUnit
}

func calculateRecommendedReportItemLengthFloat(columnName string, results []*stats.BenchmarkResult, getField func(result *stats.BenchmarkResult) float64) int {
	length := stringLength(columnName)
	for _, result := range results {
		value := getField(result)
		length = max(length, stringLength(fmt.Sprintf("%0.3f", value)))
	}

	return length
}
