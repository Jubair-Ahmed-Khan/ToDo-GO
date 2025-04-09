package utils

// ANSI escape codes for coloring text
const (
	ColorRed     = "\033[31m"
	ColorGreen   = "\033[32m"
	ColorYellow  = "\033[33m"
	ColorBlue    = "\033[34m"
	ColorMagenta = "\033[35m"
	ColorReset   = "\033[0m"
)

// function to colorize text
func Colorize(text string, colorCode string) string {
	return colorCode + text + ColorReset
}
