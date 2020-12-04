package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type ValidationFunc func(f Field) bool

func ValidDigits(n, lo, hi int) ValidationFunc {
	return func(f Field) bool {
		if len(f.Val) < n {
			return false
		}
		n, err := strconv.Atoi(f.Val)
		if err != nil {
			return false
		}
		if n < lo || n > hi {
			return false
		}
		return true
	}
}

func ValidHeight() ValidationFunc {
	re := regexp.MustCompile(`^([0-9]+)(in|cm)$`)
	return func(f Field) bool {
		match := re.FindStringSubmatch(f.Val)
		if match == nil {
			return false
		}
		n, err := strconv.Atoi(match[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "error parsing height %s: %s", match[0], err)
			return false
		}
		switch match[2] {
		case "cm":
			if n < 150 || n > 193 {
				return false
			}
			return true
		case "in":
			if n < 59 || n > 76 {
				return false
			}
			return true
		default:
			return false
		}
	}
}

func ValidColor() ValidationFunc {
	re := regexp.MustCompile(`^#([0-9a-f]{6})$`)
	return func(f Field) bool {
		return re.MatchString(f.Val)
	}
}

func ValidEyeColor() ValidationFunc {
	validColors := []string{
		"amb", "blu", "brn", "gry", "grn", "hzl", "oth",
	}
	return func(f Field) bool {
		color := strings.ToLower(f.Val)
		for _, validColor := range validColors {
			if color == validColor {
				return true
			}
		}
		return false
	}
}

func ValidPasswordId() ValidationFunc {
	re := regexp.MustCompile(`^([0-9]{9})$`)
	return func(f Field) bool {
		return re.MatchString(f.Val)
	}
}

func AlwaysValid() ValidationFunc {
	return func(f Field) bool {
		return true
	}
}

type Field struct {
	Key string
	Val string
	Opt bool
	VFn ValidationFunc
}

func (f Field) String() string {
	return fmt.Sprintf("%s:%s (Optional: %t", f.Key, f.Val, f.Opt)
}

func (f Field) Valid() bool {
	return f.VFn(f)
}

type Passport struct {
	fields map[string]*Field
}

func NewPassport() *Passport {
	fields := map[string]*Field{
		"byr": &Field{Key: "byr", VFn: ValidDigits(4, 1920, 2002)},
		"iyr": &Field{Key: "iyr", VFn: ValidDigits(4, 2010, 2020)},
		"eyr": &Field{Key: "eyr", VFn: ValidDigits(4, 2020, 2030)},
		"hgt": &Field{Key: "hgt", VFn: ValidHeight()},
		"hcl": &Field{Key: "hcl", VFn: ValidColor()},
		"ecl": &Field{Key: "ecl", VFn: ValidEyeColor()},
		"pid": &Field{Key: "pid", VFn: ValidPasswordId()},
		"cid": &Field{Key: "cid", VFn: AlwaysValid(), Opt: true},
	}

	return &Passport{fields}
}

func (p *Passport) String() string {
	var fields []string
	for _, f := range p.fields {
		fields = append(fields, fmt.Sprintf("%s:%s", f.Key, f.Val))
	}
	return strings.Join(fields, " ")
}

func (p *Passport) SetField(key, val string) {
	p.fields[key].Val = val
}

func (p *Passport) Valid() bool {
	for _, f := range p.fields {
		if f.Val == "" && f.Opt {
			continue
		}
		if !f.Valid() {
			fmt.Fprintf(
				os.Stderr,
				"invalid passport: %q\n field: %q\n",
				p, f,
			)
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
