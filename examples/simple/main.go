package main

import (
	"log"

	"github.com/cubixle/l1"
)

func main() {
	r, err := l1.NewRunner(
		l1.WithTarget("http://google.com"),
		l1.WithRunFunc(func(target string) error {
			return nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}
	err = r.Start()
	if err != nil {
		log.Fatal(err)
	}
}
