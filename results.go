package benchy

import (
	"os"
	"testing"

	"github.com/smarty/benchy/stats"
)

// ReadResultsFromFile opens the provided file and reads the collection of
// benchmark results from it.
//
// Parameters:
//   - tb is the current testing handle, whether that's *testing.T or *testing.B.
//   - filename is the location of the file to open and read from.
//
// Returns:
//   - results is the collection of benchmark results that were read from file.
//   - err is nil on a successful operation. If unsuccessful, error contains
//     whatever error was returned from attempting to open the file or reading
//     from the file.
func ReadResultsFromFile(tb testing.TB, filename string) (results *stats.BenchmarkResults, err error) {
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return nil, err
	}

	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	results = stats.NewBenchmarkResults(tb)
	_, err = results.ReadFrom(file)
	return results, err
}

// WriteResultsToFile creates the provided file and writes the collection of
// benchmark results to it. This function will overwrite rather than append.
//
// Parameters:
//   - filename is the location that the file will be written.
//   - results is the collection of results to write to file.
//
// Returns:
//   - err is `nil` on a successful operation. On failure to write to file,
//     err contains the error that was returned from either trying to create
//     the file or from writing to the file.
func WriteResultsToFile(filename string, results *stats.BenchmarkResults) (err error) {
	var file *os.File
	if file, err = os.Create(filename); err != nil {
		return err
	}

	defer func() {
		if err = file.Close(); err != nil {
			panic(err)
		}
	}()

	_, err = results.WriteTo(file)
	return err
}
