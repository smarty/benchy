package params

import (
	"fmt"
	"testing"

	"github.com/smarty/benchy/options"
)

func Test_SelectSampleCount_FromCLI(t *testing.T) {
	expected := 10
	args := []string{"-test.samples", fmt.Sprintf("%d", expected)}

	actual := SelectSampleCount(0, options.FullMetrics, args)

	if actual != expected {
		t.Errorf("SelectSampleCount() is %v, want %v", actual, expected)
	}
}

func Test_SelectSampleCount_FromCLI_0(t *testing.T) {
	expected := fullMetricsDefault
	args := []string{"samples", "0"}

	actual := SelectSampleCount(0, options.FullMetrics, args)

	if actual != expected {
		t.Errorf("SelectSampleCount() is %v, want %v", actual, expected)
	}
}

func Test_SelectSampleCount_FromCLI_Incomplete(t *testing.T) {
	expected := fullMetricsDefault
	args := []string{"-samples"}

	actual := SelectSampleCount(0, options.FullMetrics, args)

	if actual != expected {
		t.Errorf("SelectSampleCount() is %v, want %v", actual, expected)
	}
}

func Test_SelectSampleCount_FromInput(t *testing.T) {
	expected := 10
	var args []string

	actual := SelectSampleCount(expected, options.FullMetrics, args)

	if actual != expected {
		t.Errorf("SelectSampleCount() is %v, want %v", actual, expected)
	}
}
