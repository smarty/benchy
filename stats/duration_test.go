package stats

import (
	"testing"
)

func TestDuration_Unit(t *testing.T) {
	type valueExpected struct {
		Value    Duration
		Expected string
	}

	tests := []valueExpected{
		{Value: Duration(1), Expected: nanosecondsUnit},
		{Value: Duration(100), Expected: microsecondsUnit},
		{Value: Duration(1_000), Expected: microsecondsUnit},
		{Value: Duration(100_000), Expected: millisecondsUnit},
		{Value: Duration(1_000_000), Expected: millisecondsUnit},
		{Value: Duration(100_000_000), Expected: secondsUnit},
		{Value: Duration(1_000_000_000), Expected: secondsUnit},
	}

	for iTest, test := range tests {
		actual := test.Value.Unit()
		if actual != test.Expected {
			t.Errorf("test %d failed: expected %s but got %s", iTest, test.Expected, actual)
		}
	}
}

func TestDuration_RenderLength(t *testing.T) {
	type valueExpected struct {
		Value    Duration
		Expected int
	}

	tests := []valueExpected{
		{Value: Duration(-0.3), Expected: 9},
		{Value: Duration(0.3), Expected: 9},
		{Value: Duration(1), Expected: 9},
		{Value: Duration(100), Expected: 9},
		{Value: Duration(1_000), Expected: 9},
		{Value: Duration(10_000), Expected: 10},
		{Value: Duration(100_000), Expected: 9},
		{Value: Duration(1_000_000), Expected: 9},
		{Value: Duration(100_000_000), Expected: 8},
		{Value: Duration(1_000_000_000), Expected: 8},
	}

	for iTest, test := range tests {
		actual := test.Value.RenderLength()
		if actual != test.Expected {
			t.Errorf("test %d failed: expected %d but got %d", iTest, test.Expected, actual)
		}
	}
}

func TestDuration_RenderLengthAsUnit(t *testing.T) {
	type valueExpected struct {
		Value    Duration
		Unit     string
		Expected int
	}

	tests := []valueExpected{
		{Value: Duration(-0.3), Unit: nanosecondsUnit, Expected: 9},
		{Value: Duration(-0.3), Unit: microsecondsUnit, Expected: 9},
		{Value: Duration(-0.3), Unit: millisecondsUnit, Expected: 9},
		{Value: Duration(-0.3), Unit: secondsUnit, Expected: 8},

		{Value: Duration(0.3), Unit: nanosecondsUnit, Expected: 9},
		{Value: Duration(0.3), Unit: microsecondsUnit, Expected: 9},
		{Value: Duration(0.3), Unit: millisecondsUnit, Expected: 9},
		{Value: Duration(0.3), Unit: secondsUnit, Expected: 8},

		{Value: Duration(1_000), Unit: nanosecondsUnit, Expected: 12},
		{Value: Duration(1_000), Unit: microsecondsUnit, Expected: 9},
		{Value: Duration(1_000), Unit: millisecondsUnit, Expected: 9},
		{Value: Duration(1_000), Unit: secondsUnit, Expected: 8},

		{Value: Duration(1_000_000), Unit: nanosecondsUnit, Expected: 15},
		{Value: Duration(1_000_000), Unit: microsecondsUnit, Expected: 12},
		{Value: Duration(1_000_000), Unit: millisecondsUnit, Expected: 9},
		{Value: Duration(1_000_000), Unit: secondsUnit, Expected: 8},

		{Value: Duration(1_000_000_000), Unit: nanosecondsUnit, Expected: 18},
		{Value: Duration(1_000_000_000), Unit: microsecondsUnit, Expected: 15},
		{Value: Duration(1_000_000_000), Unit: millisecondsUnit, Expected: 12},
		{Value: Duration(1_000_000_000), Unit: secondsUnit, Expected: 8},
	}

	for iTest, test := range tests {
		actual := test.Value.RenderLengthAsUnit(test.Unit)
		if actual != test.Expected {
			t.Errorf("test %d failed: expected %d but got %d", iTest, test.Expected, actual)
		}
	}
}

func TestDuration_Render(t *testing.T) {
	type valueExpected struct {
		Value    Duration
		Expected string
	}

	tests := []valueExpected{
		{Value: Duration(-0.3), Expected: "-0.300 ns"},
		{Value: Duration(0.3), Expected: " 0.300 ns"},
		{Value: Duration(1_000), Expected: " 1.000 µs"},
		{Value: Duration(1_000_000), Expected: " 1.000 ms"},
		{Value: Duration(1_000_000_000), Expected: " 1.000 s"},
	}

	for iTest, test := range tests {
		actual := test.Value.Render()
		if actual != test.Expected {
			t.Errorf("test %d failed: expected %s but got %s", iTest, test.Expected, actual)
		}
	}
}

func TestDuration_RenderAsUnit(t *testing.T) {
	type valueExpected struct {
		Value    Duration
		Unit     string
		Expected string
	}

	tests := []valueExpected{
		{Value: Duration(-0.3), Unit: nanosecondsUnit, Expected: "-0.300 ns"},
		{Value: Duration(-0.3), Unit: microsecondsUnit, Expected: "-0.000 µs"},
		{Value: Duration(-0.3), Unit: millisecondsUnit, Expected: "-0.000 ms"},
		{Value: Duration(-0.3), Unit: secondsUnit, Expected: "-0.000 s"},

		{Value: Duration(0.3), Unit: nanosecondsUnit, Expected: " 0.300 ns"},
		{Value: Duration(0.3), Unit: microsecondsUnit, Expected: " 0.000 µs"},
		{Value: Duration(0.3), Unit: millisecondsUnit, Expected: " 0.000 ms"},
		{Value: Duration(0.3), Unit: secondsUnit, Expected: " 0.000 s"},

		{Value: Duration(1_000), Unit: nanosecondsUnit, Expected: " 1000.000 ns"},
		{Value: Duration(1_000), Unit: microsecondsUnit, Expected: " 1.000 µs"},
		{Value: Duration(1_000), Unit: millisecondsUnit, Expected: " 0.001 ms"},
		{Value: Duration(1_000), Unit: secondsUnit, Expected: " 0.000 s"},

		{Value: Duration(1_000_000), Unit: nanosecondsUnit, Expected: " 1000000.000 ns"},
		{Value: Duration(1_000_000), Unit: microsecondsUnit, Expected: " 1000.000 µs"},
		{Value: Duration(1_000_000), Unit: millisecondsUnit, Expected: " 1.000 ms"},
		{Value: Duration(1_000_000), Unit: secondsUnit, Expected: " 0.001 s"},

		{Value: Duration(1_000_000_000), Unit: nanosecondsUnit, Expected: " 1000000000.000 ns"},
		{Value: Duration(1_000_000_000), Unit: microsecondsUnit, Expected: " 1000000.000 µs"},
		{Value: Duration(1_000_000_000), Unit: millisecondsUnit, Expected: " 1000.000 ms"},
		{Value: Duration(1_000_000_000), Unit: secondsUnit, Expected: " 1.000 s"},
	}

	for iTest, test := range tests {
		actual := test.Value.RenderWithUnit(test.Unit)
		if actual != test.Expected {
			t.Errorf("test %d failed: expected %s but got %s", iTest, test.Expected, actual)
		}
	}
}

func TestSmallestUnit(t *testing.T) {
	type valueExpected struct {
		Value    []string
		Expected string
	}

	tests := []valueExpected{
		{Value: []string{}, Expected: secondsUnit},
		{Value: []string{secondsUnit}, Expected: secondsUnit},
		{Value: []string{secondsUnit, secondsUnit}, Expected: secondsUnit},
		{Value: []string{secondsUnit, nanosecondsUnit}, Expected: nanosecondsUnit},
		{Value: []string{microsecondsUnit, nanosecondsUnit}, Expected: nanosecondsUnit},
		{Value: []string{microsecondsUnit, millisecondsUnit}, Expected: microsecondsUnit},
	}

	for iTest, test := range tests {
		actual := SmallestUnit(test.Value...)
		if actual != test.Expected {
			t.Errorf("test %d failed: expected %s but got %s", iTest, test.Expected, actual)
		}
	}
}
