package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/prologic/aoc"
)

func readBoardingpassesFromReader(r io.Reader) (bps aoc.Boardingpasses, err error) {
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

	bpsMap := make([]*aoc.Boardingpass, 1024)
	for _, bp := range bps {
		bpsMap[bp.SeatID()] = &bp
	}

	fmt.Fprintf(os.Stderr, "%d boarding passes found\n", len(bps))

	for i := 1; i < 1023; i++ {
		bp := bpsMap[i]
		if bp == nil {
			if bpsMap[i-1] != nil && bpsMap[i+1] != nil {
				fmt.Printf("%d\n", i)
				break
			}
		}
	}
}
