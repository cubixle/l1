package l1

type Results struct {
	Target  string
	Count   int
	Results []Result
}

type Result struct {
	// CompletedIn is in seconds
	CompletedIn int
	Error       error
	StatusCode  int
}

func (r *Results) RequestsPerMin() int {
	totalCompletedIn := 0
	for _, res := range r.Results {
		totalCompletedIn += res.CompletedIn
	}
	return totalCompletedIn / r.Count
}

func (r *Results) AvgCompletionTime() int {
	return 0
}
