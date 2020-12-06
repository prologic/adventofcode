package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	sum, g := 0, 0
	qs := make(map[rune]int)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			for _, n := range qs {
				if n == g {
					sum++
				}
			}
			g = 0
			qs = make(map[rune]int)
			continue
		}
		for _, c := range line {
			qs[c]++
		}
		g++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error reading standard input: %s", err)
		os.Exit(2)
	}
	for _, n := range qs {
		if n == g {
			sum++
		}
	}

	fmt.Printf("%d\n", sum)
}
