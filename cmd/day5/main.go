package main

import (
	"fmt"
	"os"

	"github.com/prologic/aoc"
)

func main() {
	s := os.Args[1]

	bp, err := aoc.ParseBoardingpass(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error parsing boarding pass: %s", err)
		os.Exit(2)
	}

	fmt.Printf("%s\n", bp)
}
