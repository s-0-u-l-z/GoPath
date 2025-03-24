package utils

import "github.com/fatih/color"

func StatusColor(status int) *color.Color {
	switch {
	case status >= 200 && status < 300:
		return color.New(color.FgCyan) // 200 OK -> Cyan
	case status >= 300 && status < 400:
		return color.New(color.FgBlue) // 302 Redirect -> Light Blue
	case status >= 400 && status < 500:
		return color.New(color.FgYellow) // 404 Not Found -> Yellow
	case status >= 500:
		return color.New(color.FgRed) // 500 Server Error -> Red
	default:
		return color.New(color.FgWhite)
	}
}

