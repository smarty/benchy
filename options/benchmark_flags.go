package options

// BenchmarkFlag is used to indicate options on a benchmark.
type BenchmarkFlag int

const (
	// Long indicates that a benchmark should not be run in short mode.
	Long BenchmarkFlag = 1 << iota

	// OverheadSampling explicitly turns on overhead sampling for a benchmark.
	// Overhead sampling is implicitly on when a setup or cleanup is registered
	// for the benchmark.
	//
	// Default is off.
	OverheadSampling

	// PProfCPU turns on CPU profiling for a benchmark. If PProf is being used
	// from a higher level (for example, from the CLI), then the benchmark will
	// fail rather than try to take over the PProf that is already running.
	//
	// When turned on, PProf files will be saved in a folder call "workspace" in
	// the current working directory.
	PProfCPU
)

// Contains determines if all the indicated flags are set in this flags value.
//
// Parameters:
//   - flags are the values to test for in this flag.
//
// Returns:
//   - found is `true` if all `flags` were found in this flag.
func (this *BenchmarkFlag) Contains(flags ...BenchmarkFlag) (found bool) {
	for _, flag := range flags {
		if flag&*this != flag {
			return false
		}
	}

	return true
}

// Set adds all the indicated flags from this flags value.
//
// Parameters:
//   - flags are the values to set in this flag.
func (this *BenchmarkFlag) Set(flags ...BenchmarkFlag) {
	for _, flag := range flags {
		*this |= flag
	}
}

// Clear removes all the indicated flags from this flags value.
//
// Parameters:
//   - flags are the values to unset in this flag.
func (this *BenchmarkFlag) Clear(flags ...BenchmarkFlag) {
	for _, flag := range flags {
		*this &= ^flag
	}
}
