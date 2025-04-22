package scanner

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func MakeRequest(url string, args Arguments) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	status := resp.StatusCode

	var printer func(a ...interface{}) string

	switch {
	case status >= 200 && status < 300:
		printer = color.New(color.FgGreen).SprintFunc()
	case status >= 300 && status < 400:
		printer = color.New(color.FgCyan).SprintFunc()
	case status >= 400 && status < 500:
		printer = color.New(color.FgYellow).SprintFunc()
	case status >= 500:
		printer = color.New(color.FgHiRed).SprintFunc()
	default:
		printer = color.New(color.FgWhite).SprintFunc()
	}

	fmt.Printf("%s %s\n", printer(fmt.Sprintf("[%d]", status)), url)
}
