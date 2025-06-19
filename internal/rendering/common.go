package rendering

const (
	ansi_red     = "\033[91m"
	ansi_green   = "\033[92m"
	ansi_yellow  = "\033[93m"
	ansi_blue    = "\033[94m"
	ansi_magenta = "\033[95m"
	ansi_cyan    = "\033[96m"
	ansi_white   = "\033[97m"
	ansi_reset   = "\033[0m"
)

func padRight(value string, totalLength int, paddingRune rune) string {
	padding := make([]rune, max(0, totalLength-stringLength(value)))
	if len(padding) == 0 {
		return value
	}

	for i := range padding {
		padding[i] = paddingRune
	}

	return value + string(padding)
}

func padLeft(value string, totalLength int, paddingRune rune) string {
	padding := make([]rune, max(0, totalLength-stringLength(value)))
	if len(padding) == 0 {
		return value
	}

	for i := range padding {
		padding[i] = paddingRune
	}

	return string(padding) + value
}

func stringLength(value string) int {
	length := 0
	for range value {
		length++
	}

	return length
}
