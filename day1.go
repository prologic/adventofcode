package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func readIntsFromReader(r io.Reader) (xs []int, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		x, err := strconv.Atoi(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error converting %s to int: %s", line, err)
			continue
		}
		xs = append(xs, x)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading standard input: %w", err)
	}

	return
}

func findPairSum(xs []int, s int) (x, y int, err error) {
	for _, x = range xs {
		for _, y = range xs {
			if x+y == s {
				return
			}
		}
	}

	return 0, 0, fmt.Errorf("error: no pair found summing to %d", s)
}

func findTriSum(xs []int, s int) (x, y, z int, err error) {
	for _, x = range xs {
		for _, y = range xs {
			for _, z = range xs {
				if x+y+z == s {
					return
				}
			}
		}
	}

	return 0, 0, 0, fmt.Errorf("error: no tri found summing to %d", s)
}

func main() {
	xs, err := readIntsFromReader(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %s", err)
		os.Exit(2)
	}

	x, y, err := findPairSum(xs, 2020)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error finding pair sum: %s", err)
		os.Exit(2)
	}

	fmt.Printf("%d\n", x*y)

	x, y, z, err := findTriSum(xs, 2020)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error finding tri sum: %s", err)
		os.Exit(2)
	}

	fmt.Printf("%d\n", x*y*z)

}
