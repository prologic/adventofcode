package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/prologic/aoc"
)

func readBoardingpassesFromReader(r io.Reader) (bps []aoc.Boardingpass, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		bp, err := aoc.ParseBoardingpass(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing boarding pass %s: %s", line, err)
			continue
		}

		bps = append(bps, bp)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading standard input: %w", err)
	}

	return
}

func main() {
	bps, err := readBoardingpassesFromReader(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %s", err)
		os.Exit(2)
	}

	fmt.Fprintf(os.Stderr, "%d boarding passes found\n", len(bps))

	var highestSeatID int

	for _, bp := range bps {
		fmt.Fprintf(os.Stderr, "%s\n", bp)
		if bp.SeatID() > highestSeatID {
			highestSeatID = bp.SeatID()
		}
	}

	fmt.Printf("%d\n", highestSeatID)
}
