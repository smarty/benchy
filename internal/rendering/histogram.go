package rendering

import (
	"fmt"
	"strings"

	. "github.com/smarty/benchy/stats"
)

// Histogram renders the histogram as a series of lines which can be written out.
//
// Ansi codes are used to color the text.
func Histogram(result BenchmarkResult) []string {
	// add a line for the header and for the footer
	lines := make([]string, 0, len(result.Histogram)+2)
	lines = append(lines, fmt.Sprintf(
		"%s%s: %s%s",
		ansi_cyan,
		result.Name,
		ansi_reset,
		generateModalityString(result)))

	start, step := calculateStepping(result.Histogram, result.Min, result.Max)
	low := start
	high := start
	truncatedPrevious := false
	for bucketNumber, numberInBucket := range result.Histogram {
		low = high
		high = low + step
		if shouldSkipBucket(result, numberInBucket, bucketNumber) {
			if !truncatedPrevious {
				lines = append(lines, fmt.Sprintf(
					"%s%s...%s",
					ansi_cyan,
					padRight("", 25, ' '),
					ansi_reset))

				truncatedPrevious = true
			}

			continue
		}

		truncatedPrevious = false
		number := ""
		if numberInBucket != 0 {
			number = fmt.Sprintf("%d", numberInBucket)
		}

		lines = append(lines, fmt.Sprintf(
			"%s%s |%s %s%s",
			ansi_cyan,
			padLeft(fmt.Sprintf("%s - %s", low.Render(), high.Render()), 25, ' '),
			renderBucketSamples(result, bucketNumber, low, high),
			number,
			ansi_reset))
	}

	if len(result.Outliers) > 0 {
		lines = append(lines, listOutliers(result))
	}

	return lines
}

func shouldSkipBucket(result BenchmarkResult, numberInBucket int, bucketNumber int) bool {
	return numberInBucket == 0 && result.Histogram[bucketNumber-1] == 0 && result.Histogram[bucketNumber+1] == 0
}

func renderBucketSamples(result BenchmarkResult, bucketNumber int, low Duration, high Duration) string {
	outliersInBucket := 0
	for _, outlier := range result.Outliers {
		if outlier >= low-0.0001 && outlier <= high+0.0001 {
			outliersInBucket++
		}
	}

	numberInBucket := result.Histogram[bucketNumber]
	outliers := padRight("", outliersInBucket, 'X')
	regular := padRight("", numberInBucket-outliersInBucket, 'X')
	blanks := padRight("", len(result.Samples)-numberInBucket, ' ')
	if bucketNumber < len(result.Histogram)/2 {
		return fmt.Sprintf("%s%s%s%s%s", ansi_yellow, outliers, ansi_cyan, regular, blanks)
	}

	return fmt.Sprintf("%s%s%s%s%s", regular, ansi_yellow, outliers, ansi_cyan, blanks)
}

func calculateStepping(histogram []int, minimum Duration, maximum Duration) (start Duration, step Duration) {
	start = minimum
	step = (maximum - minimum) / Duration(len(histogram))
	return start, step
}

func generateModalityString(result BenchmarkResult) string {
	if result.Modality == 1 {
		return fmt.Sprintf("%suni-modal%s", ansi_cyan, ansi_reset)
	}

	category := "multi"
	switch result.Modality {
	case 0:
		category = "non"
	case 2:
		category = "bi"
	}

	return fmt.Sprintf("%s%s-modal%s", ansi_yellow, category, ansi_reset)
}

func listOutliers(result BenchmarkResult) string {
	sb := strings.Builder{}
	sb.WriteString(ansi_yellow)

	preamble := "outliers were"
	if len(result.Outliers) == 1 {
		preamble = "outlier was"
	}

	sb.WriteString(fmt.Sprintf("%d %s removed: [", len(result.Outliers), preamble))
	for iOutlier, outlier := range result.Outliers {
		if iOutlier != 0 {
			sb.WriteString(",")
		}

		sb.WriteString(outlier.Render())
	}

	sb.WriteString(" ]")
	sb.WriteString(ansi_reset)
	return sb.String()
}
