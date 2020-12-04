package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Field struct {
	Key string
	Val string
	Opt bool
}

type Passport struct {
	fields map[string]*Field
}

func NewPassport() *Passport {
	fields := map[string]*Field{
		"byr": &Field{Key: "byr"},
		"iyr": &Field{Key: "iyr"},
		"eyr": &Field{Key: "eyr"},
		"hgt": &Field{Key: "hgt"},
		"hcl": &Field{Key: "hcl"},
		"ecl": &Field{Key: "ecl"},
		"pid": &Field{Key: "pid"},
		"cid": &Field{Key: "cid", Opt: true},
	}

	return &Passport{fields}
}

func (p *Passport) SetField(key, val string) {
	p.fields[key].Val = val
}

func (p *Passport) Valid() bool {
	for _, f := range p.fields {
		if f.Val == "" && !f.Opt {
			return false
		}
	}
	return true
}

func readPassportsFromReader(r io.Reader) (ps []*Passport, err error) {
	scanner := bufio.NewScanner(r)
	aPassport := NewPassport()
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			ps = append(ps, aPassport)
			aPassport = NewPassport()
			continue
		}
		for _, pair := range strings.Split(line, " ") {
			tokens := strings.Split(pair, ":")
			if len(tokens) > 2 {
				fmt.Fprintf(
					os.Stderr,
					"error parsing passport field expected 2 tokens got %d",
					len(tokens),
				)
				continue
			}
			key := strings.ToLower(tokens[0])
			val := tokens[1]
			aPassport.SetField(key, val)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading standard input: %w", err)
	}
	// Lass passport
	ps = append(ps, aPassport)

	return
}

func main() {
	ps, err := readPassportsFromReader(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %s", err)
		os.Exit(2)
	}

	fmt.Fprintf(os.Stderr, "%d passports found\n", len(ps))

	var valid int
	for _, p := range ps {
		if p.Valid() {
			valid++
		}
	}

	fmt.Printf("%d\n", valid)
}
