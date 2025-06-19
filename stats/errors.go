package stats

import "fmt"

var (
	CouldNotFindResultError = fmt.Errorf("could not find benchmark result")
)

func generateBenchmarkNotFoundError(benchmarkName string) error {
	return fmt.Errorf(
		"%w: benchmark results for \"%s\" could not be found, ensure that you spelled it correctly",
		CouldNotFindResultError,
		benchmarkName)
}
