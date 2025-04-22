package scanner

import (
	"net/http"
	"time"
	
)

var client = &http.client{
	Timeout:5 * time.Second,
	Transport: &http.Transport{
		MaxIdleConns:	 100,
		MaxIdleConnsPerHost: 100,
		IdleConnTimeout:     90 * time.Second,
	},
}
