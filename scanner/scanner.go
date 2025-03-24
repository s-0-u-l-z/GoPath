package scanner

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"

	"github.com/akamensky/argparse"
)

type Arguments struct {
	URL              string
	URLsFile         string
	Wordlist         string
	Extensions       string
	Threads          int
	Recursive        bool
	MaxRecursion     int
	IncludeStatus    string
	ExcludeStatus    string
	Timeout          int
	Proxy            string
	UserAgent        string
	OutputFile       string
	FollowRedirects  bool
	Auth             string
	Cookie           string
	RandomAgent      bool
	QuietMode        bool
	ShowVersion      bool
}

func ParseArguments() Arguments {
	parser := argparse.NewParser("GoPath", "Fast web directory scanner in Go")

	// General options
	showVersion := parser.Flag("", "version", &argparse.Options{Help: "Show program's version number and exit"})
	help := parser.Flag("h", "help", &argparse.Options{Help: "Show this help message and exit"})

	// Target options
	url := parser.String("u", "url", &argparse.Options{Required: true, Help: "Target URL"})
	urlsFile := parser.String("l", "urls-file", &argparse.Options{Help: "File containing list of URLs to scan"})
	wordlist := parser.String("w", "wordlist", &argparse.Options{Default: "wordlist.txt", Help: "Wordlist file"})

	// Extensions & Recursive scanning
	extensions := parser.String("e", "extensions", &argparse.Options{Help: "Extensions to scan (comma-separated)"})
	recursive := parser.Flag("r", "recursive", &argparse.Options{Help: "Enable recursive scanning"})
	maxRecursion := parser.Int("R", "max-recursion-depth", &argparse.Options{Default: 3, Help: "Maximum recursion depth"})

	// Filtering results
	includeStatus := parser.String("i", "include-status", &argparse.Options{Help: "Include specific status codes (e.g. 200,300-399)"})
	excludeStatus := parser.String("x", "exclude-status", &argparse.Options{Help: "Exclude specific status codes (e.g. 404,500-599)"})

	// Performance settings
	threads := parser.Int("t", "threads", &argparse.Options{Default: 10, Help: "Number of concurrent threads"})
	timeout := parser.Int("", "timeout", &argparse.Options{Default: 5, Help: "Request timeout in seconds"})

	// Network options
	proxy := parser.String("p", "proxy", &argparse.Options{Help: "Use a proxy (http://proxy:port)"})
	userAgent := parser.String("", "user-agent", &argparse.Options{Help: "Custom User-Agent"})
	randomAgent := parser.Flag("", "random-agent", &argparse.Options{Help: "Use a random User-Agent for each request"})

	// Output settings
	outputFile := parser.String("o", "output", &argparse.Options{Help: "Output results to a file"})
	quietMode := parser.Flag("q", "quiet-mode", &argparse.Options{Help: "Run in quiet mode (only output results)"})

	// Authentication & Cookies
	auth := parser.String("", "auth", &argparse.Options{Help: "Authentication credential (user:password or token)"})
	cookie := parser.String("", "cookie", &argparse.Options{Help: "Send a custom Cookie header"})
	followRedirects := parser.Flag("F", "follow-redirects", &argparse.Options{Help: "Follow HTTP redirects"})

	// Parse arguments
	err := parser.Parse(os.Args)
	if err != nil || *help {
		fmt.Println(parser.Usage(err))
		os.Exit(1)
	}

	return Arguments{
		URL:             *url,
		URLsFile:        *urlsFile,
		Wordlist:        *wordlist,
		Extensions:      *extensions,
		Threads:         *threads,
		Recursive:       *recursive,
		MaxRecursion:    *maxRecursion,
		IncludeStatus:   *includeStatus,
		ExcludeStatus:   *excludeStatus,
		Timeout:         *timeout,
		Proxy:           *proxy,
		UserAgent:       *userAgent,
		RandomAgent:     *randomAgent,
		OutputFile:      *outputFile,
		QuietMode:       *quietMode,
		Auth:            *auth,
		Cookie:          *cookie,
		FollowRedirects: *followRedirects,
		ShowVersion:     *showVersion,
	}
}

func StartScan(args Arguments) {
	// Open the wordlist file
	file, err := os.Open(args.Wordlist)
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup
	sem := make(chan struct{}, args.Threads)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := scanner.Text()

		// Generate URL variations with extensions if provided
		var paths []string
		if args.Extensions == "" {
			paths = append(paths, fmt.Sprintf("%s/%s", args.URL, word))
		} else {
			for _, ext := range strings.Split(args.Extensions, ",") {
				paths = append(paths, fmt.Sprintf("%s/%s.%s", args.URL, word, ext))
			}
		}

		// Launch concurrent requests
		for _, path := range paths {
			wg.Add(1)
			go func(url string) {
				defer wg.Done()
				sem <- struct{}{}
				MakeRequest(url, args)
				<-sem
			}(path)
		}
	}

	wg.Wait()
}
