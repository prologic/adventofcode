package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sum := 0
	qs := make(map[rune]bool)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			n := 0
			for _, a := range qs {
				if a {
					n++
				}
			}
			sum += n
			fmt.Fprintf(os.Stderr, "%d\n", n)
			qs = make(map[rune]bool)
			continue
		}
		for _, c := range line {
			qs[c] = true
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading standard input: %s", err)
		os.Exit(2)
	}
	n := 0
	for _, a := range qs {
		if a {
			n++
		}
	}
	sum += n
	fmt.Fprintf(os.Stderr, "%d\n", n)

	fmt.Printf("%d\n", sum)
}
