package scanner

import (
	"net/http"
	"time"

	"github.com/s-0-u-l-z/GoPath/utils"
)

func MakeRequest(url string, args Arguments) {
	client := &http.Client{Timeout: time.Duration(args.Timeout) * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	color := utils.StatusColor(resp.StatusCode)
	color.Printf("[%d] %s\n", resp.StatusCode, url)
}
