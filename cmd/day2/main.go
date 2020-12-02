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
	c1 int
	c2 int
	ch rune
}

func (pol passwordPolicy) Check(pwd string) bool {
	var x int

	for _, ch := range pwd {
		if ch == pol.ch {
			x++
		}
	}

	if x >= pol.c1 && x <= pol.c2 {
		return true
	}

	return false
}

func (pol passwordPolicy) Check2(pwd string) bool {
	if pol.c1 > len(pwd) || pol.c2 > len(pwd) {
		return false
	}

	var X, Y bool

	if rune(pwd[pol.c1]) == pol.ch {
		X = true
	}
	if rune(pwd[pol.c2]) == pol.ch {
		Y = true
	}

	if (X || Y) && !(X && Y) {
		return true
	}

	return false
}

func parseInts(s string) (c1, c2 int, err error) {
	tokens := strings.Split(s, "-")
	if len(tokens) != 2 {
		return 0, 0, fmt.Errorf("error parsing ints, expected 2 tokens but got %d", len(tokens))
	}

	c1, err = strconv.Atoi(tokens[0])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing c1: %w", err)
	}

	c2, err = strconv.Atoi(tokens[1])
	if err != nil {
		return 0, 0, fmt.Errorf("error parsing c2: %w", err)
	}

	return
}

func parsePasswordPolicy(s string) (pol passwordPolicy, err error) {
	tokens := strings.Split(s, " ")
	if len(tokens) != 2 {
		return passwordPolicy{}, fmt.Errorf("error parsing password policy, expected 2 tokens but got %d", len(tokens))
	}

	c1, c2, err := parseInts(tokens[0])
	if err != nil {
		return passwordPolicy{}, fmt.Errorf("error parsing password policy: %w", err)
	}

	ch := []rune(tokens[1])[0]

	return passwordPolicy{
		c1: c1,
		c2: c2,
		ch: ch,
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

	// Old Password Policy
	valid := 0
	for i, pol := range pols {
		if pol.Check(pwds[i]) {
			valid++
		}
	}

	fmt.Printf("%d\n", valid)

	// New Password Policy
	valid = 0
	for i, pol := range pols {
		if pol.Check2(pwds[i]) {
			valid++
		}
	}

	fmt.Printf("%d\n", valid)
}
