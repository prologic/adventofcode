package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	min int
	max int
	ch  rune
}

func (pol passwordPolicy) Check(pwd string) bool {
	var x int

	for _, ch := range pwd {
		if ch == pol.ch {
			x++
		}
	}

	if x >= pol.min && x <= pol.max {
		return true
	}

	return false
}

func parseMinMax(s string) (min, max int, err error) {
	tokens := strings.Split(s, "-")
	if len(tokens) != 2 {
		return 0, 0, fmt.Errorf("error parsing min/max, expected 2 tokens but got %d", len(tokens))
	}

	min, err = strconv.Atoi(tokens[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing min: %w", err)
	}

	max, err = strconv.Atoi(tokens[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing max: %w", err)
	}

	return
}

func parsePasswordPolicy(s string) (pol passwordPolicy, err error) {
	tokens := strings.Split(s, " ")
	if len(tokens) != 2 {
		return passwordPolicy{}, fmt.Errorf("error parsing password policy, expected 2 tokens but got %d", len(tokens))
	}

	min, max, err := parseMinMax(tokens[0])
	if err != nil {
		return passwordPolicy{}, fmt.Errorf("error parsing password policy: %w", err)
	}

	ch := []rune(tokens[1])[0]

	return passwordPolicy{
		min: min,
		max: max,
		ch:  ch,
	}, nil
}

func readPoliciesAndPasswordsFromReader(r io.Reader) (pols []passwordPolicy, pwds []string, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		tokens := strings.Split(line, ":")
		if len(tokens) != 2 {
			fmt.Fprintf(os.Stderr, "error parsing line, expecting 2 tokens got %d", len(tokens))
			continue
		}

		pol, err := parsePasswordPolicy(tokens[0])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing password policy %s: %s", tokens[0], err)
			continue
		}

		pwd := tokens[1]

		pols = append(pols, pol)
		pwds = append(pwds, pwd)
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("error reading standard input: %w", err)
	}

	return
}

func main() {
	pols, pwds, err := readPoliciesAndPasswordsFromReader(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %s", err)
		os.Exit(2)
	}

	if len(pols) != len(pwds) {
		fmt.Fprintf(
			os.Stderr,
			"error mismatched password polocies vs. passwords %d/%d",
			len(pols), len(pwds),
		)
		os.Exit(2)
	}

	var valid int
	for i, pol := range pols {
		if pol.Check(pwds[i]) {
			valid++
		}
	}

	fmt.Printf("%d\n ", valid)
}
