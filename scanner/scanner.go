package scanner

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strings"
	"sync"

	"github.com/akamensky/argparse"
	"github.com/s-0-u-l-z/GoPath/utils"
)

type Arguments struct {
	URLsFile         string
	URL              string
	Wordlist         string
	Extensions       string
	Threads          int
	Mode             string
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
	parser := argparse.NewParser("GoPath", "Fastest web path scanner in Go")

	url := parser.String("u", "url", &argparse.Options{Help: "Target URL"})
	urlsFile := parser.String("l", "urls-file", &argparse.Options{Help: "File containing list of URLs to scan"})
	wordlist := parser.String("w", "wordlist", &argparse.Options{Default: "wordlist.txt", Help: "Wordlist file (default is embedded)"})
	extensions := parser.String("e", "extensions", &argparse.Options{Help: "Comma-separated extensions to try"})
	mode := parser.String("m", "mode", &argparse.Options{Default: "Standard", Help: "Scan mode: Standard or Fast"})

	recursive := parser.Flag("r", "recursive", &argparse.Options{Help: "Enable recursive scanning"})
	maxRecursion := parser.Int("R", "max-recursion-depth", &argparse.Options{Default: 3, Help: "Max recursion depth"})

	includeStatus := parser.String("i", "include-status", &argparse.Options{Help: "Include status codes (e.g. 200,301-399)"})
	excludeStatus := parser.String("x", "exclude-status", &argparse.Options{Help: "Exclude status codes (e.g. 404,500-599)"})

	threads := parser.Int("t", "threads", &argparse.Options{Default: 50, Help: "Number of threads"})
	timeout := parser.Int("", "timeout", &argparse.Options{Default: 5, Help: "Request timeout (seconds)"})

	proxy := parser.String("p", "proxy", &argparse.Options{Help: "Proxy URL (http://host:port)"})
	userAgent := parser.String("", "user-agent", &argparse.Options{Help: "Custom User-Agent"})
	randomAgent := parser.Flag("", "random-agent", &argparse.Options{Help: "Use random User-Agent"})

	outputFile := parser.String("o", "output", &argparse.Options{Help: "Write results to file"})
	quiet := parser.Flag("q", "quiet", &argparse.Options{Help: "Quiet mode (suppress logs)"})

	auth := parser.String("", "auth", &argparse.Options{Help: "Basic auth or token"})
	cookie := parser.String("", "cookie", &argparse.Options{Help: "Custom cookie"})
	followRedirects := parser.Flag("F", "follow-redirects", &argparse.Options{Help: "Follow redirects"})

	showVersion := parser.Flag("", "version", &argparse.Options{Help: "Show version and exit"})

	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Println(parser.Usage(err))
		os.Exit(1)
	}

	return Arguments{
		URL:             *url,
		URLsFile:        *urlsFile,
		Wordlist:        *wordlist,
		Extensions:      *extensions,
		Threads:         *threads,
		Mode:            *mode,
		Recursive:       *recursive,
		MaxRecursion:    *maxRecursion,
		IncludeStatus:   *includeStatus,
		ExcludeStatus:   *excludeStatus,
		Timeout:         *timeout,
		Proxy:           *proxy,
		UserAgent:       *userAgent,
		RandomAgent:     *randomAgent,
		OutputFile:      *outputFile,
		QuietMode:       *quiet,
		Auth:            *auth,
		Cookie:          *cookie,
		FollowRedirects: *followRedirects,
		ShowVersion:     *showVersion,
	}
}

type job struct {
	URL string
}

func ensureHTTP(url string) string {
	url = strings.TrimSpace(url)
	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return "http://" + url
	}
	return url
}

func StartScan(args Arguments) {
	var targets []string

	if args.URLsFile != "" {
		file, err := os.Open(args.URLsFile)
		if err != nil {
			fmt.Println("Error opening URLs file:", err)
			return
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			url := ensureHTTP(scanner.Text())
			targets = append(targets, url)
		}
	} else if args.URL != "" {
		targets = append(targets, ensureHTTP(args.URL))
	} else {
		fmt.Println("Error: No target URL provided. Use -u for a single URL or -l for a URL list.")
		return
	}

	var wordlist []string
	if args.Wordlist == "" || args.Wordlist == "wordlist.txt" {
		wordlist = utils.Wordlist
	} else {
		data, err := os.ReadFile(args.Wordlist)
		if err != nil {
			fmt.Printf("Failed to load custom wordlist: %v\n", err)
			return
		}
		wordlist = strings.Split(string(data), "\n")
	}

	if strings.ToLower(args.Mode) == "fast" {
		args.Threads = runtime.NumCPU() * 4
	}

	jobs := make(chan job, args.Threads*2)
	var wg sync.WaitGroup

	for i := 0; i < args.Threads; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := range jobs {
				MakeRequest(j.URL, args)
			}
		}()
	}

	for _, word := range wordlist {
		word = strings.TrimSpace(word)
		if word == "" {
			continue
		}
		for _, target := range targets {
			jobs <- job{URL: fmt.Sprintf("%s/%s", target, word)}
		}
	}

	close(jobs)
	wg.Wait()
}
