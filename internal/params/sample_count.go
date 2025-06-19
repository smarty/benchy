package params

import (
	"flag"
	"strconv"
	"strings"

	"github.com/smarty/benchy/options"
)

const (
	fastDefault        = 3
	mediumDefault      = 10
	fullMetricsDefault = 25
)

var _ = flag.Int("test.samples", 0, "Number of benchmark samples to run and collect.")

// SelectSampleCount looks for a user-defined sample count from the input `args` first. Then looks at `input`.
// Finally, if no user-defined sample count is defined, defaults is defined by profile.
//
// Minimum is 1.
func SelectSampleCount(input int, profile options.BenchmarkProfile, args []string) int {
	for iArgument, argument := range args {
		if iArgument == len(args)-1 {
			break
		}

		if !strings.EqualFold(argument, "-test.samples") {
			continue
		}

		nextArgument := args[iArgument+1]
		cliValue, err := strconv.Atoi(nextArgument)
		if err != nil {
			break
		}

		if cliValue < 1 {
			break
		}

		return cliValue
	}

	if input == 0 {
		switch profile {
		case options.Fast:
			return fastDefault

		case options.Medium:
			return mediumDefault

		default:
			return fullMetricsDefault
		}
	}

	return input
}
