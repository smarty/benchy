package assertions

import (
	"fmt"
	"runtime"
	"strings"
)

var (
	AssertionFailedError     = fmt.Errorf("assertion failed")
	NotEnoughBenchmarksError = fmt.Errorf("not enough benchmarks")
)

func generateNoRightHandError(leftName string) error {
	return generateError(
		"expects at least 1 right-hand benchmark to compare \"%s\" against",
		NotEnoughBenchmarksError,
		leftName)
}

func generateError(format string, innerError error, data ...any) error {
	functionName := getCallingFunctionName()

	message := fmt.Sprintf(format, data...)
	return fmt.Errorf("%w: \"%s\" %s", innerError, functionName, message)
}

func getCallingFunctionName() string {
	pc, _, _, _ := runtime.Caller(2)
	nameParts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	functionName := nameParts[len(nameParts)-1]
	return functionName
}
