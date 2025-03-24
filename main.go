package main

import (
	"fmt"
	"os"

	"github.com/s-0-u-l-z/GoPath/scanner"
	"github.com/s-0-u-l-z/GoPath/utils"
)

func main() {
	// Print the GoPath banner
	utils.PrintBanner("1.0.0")

	// Parse CLI arguments
	args := scanner.ParseArguments()

	// Handle --version flag
	if args.ShowVersion {
		fmt.Println("GoPath v1.0.0")
		os.Exit(0)
	}

	// Start the scanner with parsed arguments
	scanner.StartScan(args)
}
