package main

import (
	"log"

	"github.com/cubixle/l1"
)

func main() {
	r, err := l1.NewRunner(
		l1.WithTarget("https://remoteukjobs.com"),
		l1.WithRunFunc(l1.DefaultHTTPTester),
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
