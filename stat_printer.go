package benchy

import (
	"fmt"

	"github.com/smarty/benchy/internal/rendering"
	"github.com/smarty/benchy/stats"
)

type statPrinter interface {
	printHistogram(result *stats.BenchmarkResult, sampleCount int)
	printReportCard(results []*stats.BenchmarkResult, sampleCount int, renderingFunc rendering.ExtraRenderingFunc)
}

type activePrinter struct{}

type nullPrinter struct{}

func (this *activePrinter) printHistogram(result *stats.BenchmarkResult, sampleCount int) {
	if sampleCount >= stats.MinFullCalculation {
		fmt.Println()
		printLines(rendering.Histogram(*result))
	}
}

func (this *activePrinter) printReportCard(results []*stats.BenchmarkResult, sampleCount int, renderingFunc rendering.ExtraRenderingFunc) {
	fmt.Println()
	printLines(rendering.ReportCard(results, sampleCount, renderingFunc))
	fmt.Println()
}

func (this *nullPrinter) printHistogram(result *stats.BenchmarkResult, sampleCount int) {}

func (this *nullPrinter) printReportCard(results []*stats.BenchmarkResult, sampleCount int, renderingFunc rendering.ExtraRenderingFunc) {
}

func printLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}
