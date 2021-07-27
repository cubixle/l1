package l1

import (
	"net/http"
	"time"
)

func DefaultHTTPTester(target string) *Result {
	result := &Result{}
	client := http.Client{
		Timeout: 30 * time.Second,
	}

	req, err := http.NewRequest(http.MethodGet, target, nil)
	if err != nil {
		result.Error = err
		return result
	}

	startTime := time.Now()
	rsp, err := client.Do(req)
	if err != nil {
		result.Error = err
	}
	result.CompletedIn = time.Since(startTime).Seconds()
	result.StatusCode = rsp.StatusCode

	return result
}
