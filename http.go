package l1

import (
	"net/http"
	"time"
)

// DefaultHTTPTester is an implementation of a simple load testing func.
// It follows the F signature and returns a Result.
// Any errors that occur will be set at Result.Error.
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

	req.Header.Add("accept-encoding", "gzip, deflate, br")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:90.0) Gecko/20100101 Firefox/90.0")

	startTime := time.Now()
	rsp, err := client.Do(req)
	if err != nil {
		result.Error = err
	}
	result.CompletedIn = float64(time.Since(startTime).Milliseconds())
	result.StatusCode = rsp.StatusCode

	return result
}
