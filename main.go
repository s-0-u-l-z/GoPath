package main

import (
	"os"

	"github.com/s-0-u-l-z/GoPath/scanner"
	"github.com/s-0-u-l-z/GoPath/utils"
)

func main() {
	args := scanner.ParseArguments()
	utils.PrintBanner("1.0.0", args.Wordlist, args.Threads, args.Mode)

	if args.ShowVersion {
		os.Exit(0)
	}

	scanner.StartScan(args)
}
