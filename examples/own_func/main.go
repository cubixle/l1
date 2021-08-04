package main

import (
	"log"
	"net/http"
	"time"

	"github.com/cubixle/l1"
)

func main() {
	r, err := l1.NewRunner(
		l1.WithTarget("https://remoteukjobs.com"),
		l1.WithRunFunc(ownFunc),
		l1.WithMaxParrellConns(10),
		l1.WithMaxConns(30),
	)
	if err != nil {
		log.Fatal(err)
	}
	r.Execute()
	results := r.Results()
	results.Print()
}

func ownFunc(target string) *l1.Result {
	result := &l1.Result{}
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
