package stats

import (
	"fmt"
	"math"
)

const (
	secondsUnit      = "s"
	millisecondsUnit = "ms"
	microsecondsUnit = "µs"
	nanosecondsUnit  = "ns"

	seconds      = Duration(1_000_000_000)
	milliseconds = Duration(1_000_000)
	microseconds = Duration(1_000)

	secondsThreshold      = seconds / 10
	millisecondsThreshold = milliseconds / 10
	microsecondsThreshold = microseconds / 10

	renderedDecimalPlaces = 3
)

// Duration is an average of nanoseconds from a single sample.
type Duration float64

// RenderLength returns the length of the string when this duration is printed
// normally.
func (this *Duration) RenderLength() int {
	return this.RenderLengthAsUnit(this.Unit())
}

// RenderLengthAsUnit returns the length of the string when this duration is
// printed using the provided unit.
//
// Parameters:
//   - unit is the string representation of the unit of measurement. See units
//     below.
//
// Units:
//
//	"s" // second
//	"ms" // millisecond
//	"µs" // microsecond
//	"ns" // nanosecond
func (this *Duration) RenderLengthAsUnit(unit string) int {
	value := scaleToUnit(*this, unit)

	// minimum value of length should be 1 (to represent a 0 before the decimal place)
	length := 1
	if math.Abs(float64(value)) > 0.9 {
		// use log10+1 and round down in order to get the number of digits before the decimal place
		length = int(max(1, math.Floor(math.Log10(float64(value))+1)))
	}

	// three is used here to represent:
	// 1. the decimal place,
	// 2. space between number and unit,
	// 3. and room for negative sign if present
	length += stringLength(unit) + renderedDecimalPlaces + 3
	return length
}

// Unit returns the unit that this duration will be rendered in.
//
// Returns:
//   - unit will appear as one of the following strings: "s", "ms", "µs", "ns".
func (this *Duration) Unit() (unit string) {
	switch {
	case *this >= secondsThreshold:
		return secondsUnit

	case *this >= millisecondsThreshold:
		return millisecondsUnit

	case *this >= microsecondsThreshold:
		return microsecondsUnit

	default:
		return nanosecondsUnit
	}
}

// Render renders this duration scaled to the most fitting unit.
func (this *Duration) Render() string {
	return this.RenderWithUnit(this.Unit())
}

// RenderWithUnit renders this duration scaled to the provided unit.
func (this *Duration) RenderWithUnit(unit string) string {
	negativeSpace := ""
	if *this > 0 {
		negativeSpace = " "
	}

	return fmt.Sprintf(
		fmt.Sprintf("%s%s0.%df %s", negativeSpace, "%", renderedDecimalPlaces, unit),
		scaleToUnit(*this, unit))
}

// SmallestUnit returns the smallest unit in the provided collection.
func SmallestUnit(units ...string) string {
	unitsInOrder := []string{nanosecondsUnit, microsecondsUnit, millisecondsUnit, secondsUnit}
	index := len(unitsInOrder) - 1
	for _, unit := range units {
		for orderedIndex, orderedUnit := range unitsInOrder {
			if orderedUnit == unit {
				index = min(orderedIndex, index)
			}
		}
	}

	return unitsInOrder[index]
}

func stringLength(value string) int {
	length := 0
	for range value {
		length++
	}

	return length
}

func scaleToUnit(duration Duration, unit string) Duration {
	var divisor Duration
	switch unit {
	case secondsUnit:
		divisor = seconds

	case millisecondsUnit:
		divisor = milliseconds

	case microsecondsUnit:
		divisor = microseconds

	default:
		divisor = Duration(1)
	}

	value := duration / divisor
	return value
}
