package assertions

import (
	. "github.com/smarty/benchy/stats"
)

func IsFasterThan(left *BenchmarkResult, right ...*BenchmarkResult) error {
	if len(right) == 0 {
		return generateNoRightHandError(left.Name)
	}

	for iRight := range right {
		if left.Average >= right[iRight].Average {
			return generateError(
				"expected \"%s\" to be faster than \"%s\", but it was not",
				AssertionFailedError,
				left.Name,
				right[iRight].Name)
		}
	}

	return nil
}

func IsSlowerThan(left *BenchmarkResult, right ...*BenchmarkResult) error {
	if len(right) == 0 {
		return generateNoRightHandError(left.Name)
	}

	for iRight := range right {
		if left.Average <= right[iRight].Average {
			return generateError(
				"expected \"%s\" to be slower than \"%s\", but it was not",
				AssertionFailedError,
				left.Name,
				right[iRight].Name)
		}
	}

	return nil
}

func IsNonAllocating(left *BenchmarkResult, right ...*BenchmarkResult) error {
	if left.Allocations != 0 {
		return generateError(
			"expected \"%s\" to not allocate, but it does",
			AssertionFailedError,
			left.Name)
	}

	return nil
}
