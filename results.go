package l1

import "fmt"

type results struct {
	Target  string
	Results []*Result

	totalCompletedTime float64
	successfulCount    int
	errorCount         int
}

type Result struct {
	// CompletedIn is in seconds
	RunTime     int64
	CompletedIn float64
	Error       error
	StatusCode  int
}

func (r *results) RequestPerSecond() float64 {
	totalCompletedIn := r.CompletedTime() / 1000
	return float64(len(r.Results)) / totalCompletedIn
}

func (r *results) PeakResponseTime() {

}

func (r *results) ErrorCount() int {
	if r.errorCount > 0 {
		return r.errorCount
	}
	for _, res := range r.Results {
		if res.Error == nil {
			continue
		}
		r.errorCount++
	}
	return r.errorCount
}

func (r *results) SuccessfulCount() int {
	if r.successfulCount > 0 {
		return r.successfulCount
	}
	for _, res := range r.Results {
		if res.Error != nil {
			continue
		}
		r.successfulCount++
	}
	return r.successfulCount
}

func (r *results) AvgResponseTime() float64 {
	totalCompletedIn := r.CompletedTime()

	return totalCompletedIn / float64(len(r.Results))
}

func (r *results) CompletedTime() float64 {
	if r.totalCompletedTime > 0 {
		return r.totalCompletedTime
	}

	for _, res := range r.Results {
		if res.Error != nil {
			continue
		}
		r.totalCompletedTime += res.CompletedIn
	}
	return r.totalCompletedTime
}

func (r *results) Print() {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("| L1 load tester.")
	fmt.Println("| Default result printer.")
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("")
	fmt.Printf("Load testing %s\n", r.Target)
	fmt.Println("")
	fmt.Printf("Request per second: %.2f\n", r.RequestPerSecond())
	fmt.Printf("Average response time: %.0f ms\n", r.AvgResponseTime())
	fmt.Printf("Success count: %d\n", r.SuccessfulCount())
	fmt.Printf("Error count: %d\n", r.ErrorCount())
	fmt.Println("")
}
