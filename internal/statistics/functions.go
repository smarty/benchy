package statistics

import (
	"math"
	"sort"
)

// Sort sorts the `collection`.
func Sort[T ~float64](collection []T) {
	sort.Slice(
		collection,
		func(iLeft int, iRight int) bool { return collection[iLeft] < collection[iRight] })
}

// Average finds the single average from the `collection`.
func Average[T ~float64](collection []T) T {
	total := T(0)
	for _, value := range collection {
		total += value
	}

	return total / T(len(collection))
}

// Median finds the median from the `collection`.
func Median[T ~float64](collection []T) T {
	Sort(collection)
	return collection[len(collection)/2]
}

// MinMax finds both the minimum and maximum of the `collection`.
func MinMax[T ~float64](collection []T) (minimum T, maximum T) {
	if len(collection) == 0 {
		return 0, 0
	}

	minimum = collection[0]
	maximum = collection[0]

	for _, value := range collection {
		minimum = min(minimum, value)
		maximum = max(maximum, value)
	}

	return minimum, maximum
}

// SeparateOutliers splits the `collection` into core and outlier values.
func SeparateOutliers[T ~float64](collection []T) (core []T, outliers []T) {
	core = make([]T, 0, len(collection))
	outliers = make([]T, 0, len(collection)/2)
	quartile1, quartile3 := quartiles1and3(collection)
	iqr := quartile3 - quartile1
	lowerBound := quartile1 - (1.5 * iqr)
	upperBound := quartile3 + (1.5 * iqr)
	for _, value := range collection {
		if value < lowerBound || value > upperBound {
			outliers = append(outliers, value)
			continue
		}

		core = append(core, value)
	}

	return core, outliers
}

// StandardDeviation calculates the standard deviation of the `collection`.
func StandardDeviation[T ~float64](collection []T) T {
	average := Average(collection)
	sum := float64(0)
	for _, value := range collection {
		sum += math.Pow(float64(value-average), 2)
	}

	variance := sum / float64(len(collection)-1)
	return T(math.Sqrt(variance))
}

// StandardError calculates the standard error of the `collection`.
func StandardError[T ~float64](collection []T) T {
	return StandardDeviation(collection) / T(math.Sqrt(float64(len(collection))))
}

// ModalityFromHistogram finds the number of modes in the `histogram`.
func ModalityFromHistogram(histogram []int) int {
	if len(histogram) == 1 {
		return 1
	}

	modes := 0
	rising := true
	for i := 1; i < len(histogram); i++ {
		if histogram[i] > histogram[i-1] {
			rising = true
		} else if histogram[i] < histogram[i-1] {
			if rising {
				modes++
				rising = false
			}
		}
	}

	if rising {
		modes++
	}

	return modes
}

// Histogram builds an automatically fitted histogram for the 'collection'.
func Histogram[T ~float64](collection []T) []int {
	low := collection[0]
	high := collection[0]
	for _, value := range collection {
		low = min(low, value)
		high = max(high, value)
	}

	bins := min(max(FreedmanDiaconisBins(collection), 1), 100)
	histogram := make([]int, bins)
	interval := (high - low) / T(bins)
	for _, value := range collection {
		index := int((value - low) / interval)
		if index >= bins {
			index = bins - 1
		}

		histogram[index]++
	}

	return histogram
}

// FreedmanDiaconisBins determines the number of evenly-distributed bins for constructing a histogram from `collection`.
func FreedmanDiaconisBins[T ~float64](collection []T) int {
	iqr := InterQuartileRange(collection)
	binSize := 2 * iqr / T(math.Cbrt(float64(len(collection))))
	dataRange := collection[len(collection)-1] - collection[0]

	// in some unfortunate cases, the calculated number of bins dips below 2,
	// we ensure that the returned value never does.
	return max(2, int(math.Ceil(float64(dataRange/binSize))))
}

// InterQuartileRange calculates the iqr from `collection`.
func InterQuartileRange[T ~float64](collection []T) T {
	quartile1, quartile3 := quartiles1and3(collection)
	return quartile3 - quartile1
}

// FourSigma calculates the 4th sigma based on the input `samples`.
func FourSigma[T ~float64](samples []T, standardDeviation T) T {
	fourX := 4 * standardDeviation
	histogram := Histogram(samples)
	modality := ModalityFromHistogram(histogram)
	if modality == 0 {
		return Median(samples) + fourX
	}

	minimum, maximum := MinMax(samples)
	step := (maximum - minimum) / T(len(histogram))
	bucketNumber := highestModeBucket(histogram)
	return minimum + (T(bucketNumber) * step) + fourX
}

func quartiles1and3[T ~float64](collection []T) (quartile1 T, quartile3 T) {
	Sort(collection)
	quartile1 = collection[int(math.Ceil(float64(len(collection))/4))]
	quartile3 = collection[int(math.Floor((float64(len(collection))*3)/4))]
	return quartile1, quartile3
}

func highestModeBucket(histogram []int) int {
	for i := len(histogram) - 2; i >= 0; i-- {
		if histogram[i] < histogram[i+1] {
			return i
		}
	}

	return 0
}
