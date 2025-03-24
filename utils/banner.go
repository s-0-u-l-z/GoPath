package utils

import (
	"fmt"
	"github.com/fatih/color"
)

func PrintBanner(version string) {
	darkBlue := color.New(color.FgBlue).SprintFunc()
	lightCyan := color.New(color.FgCyan).SprintFunc()
	yellow := color.New(color.FgYellow).SprintFunc()

	fmt.Println(darkBlue(`  e88~~\           888~-_               d8   888      `))
	fmt.Println(darkBlue(` d888      e88~-_  888   \    /~~~8e  _d88__ 888-~88e `))
	fmt.Println(darkBlue(` 8888 __  d888   i 888    |       88b  888   888  888 `))
	fmt.Println(darkBlue(` 8888   | 8888   | 888   /   e88~-888  888   888  888 `))
	fmt.Println(darkBlue(` Y888   | Y888   ' 888_-~   C888  888  888   888  888 `))
	fmt.Println(darkBlue(`  "88__/   "88_-~  888       "88_-888  "88_/ 888  888 `))
	fmt.Println(lightCyan("------------------------------------------------------------"))
	fmt.Println(yellow("▶ Mode: Standard   ▶ HTTP method: GET   ▶ Threads: 25   ▶ Wordlist: Auto"))
	fmt.Println(lightCyan("------------------------------------------------------------\n"))
}

