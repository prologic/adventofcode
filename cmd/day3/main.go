package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	Open = rune('.')
	Tree = rune('#')
)

type Cell struct {
	sym rune
}

func (c Cell) Open() bool {
	return c.sym == Open
}

func (c Cell) Tree() bool {
	return c.sym == Tree
}

func (c Cell) String() string {
	return string(c.sym)
}

type Row []Cell

func (r Row) String() string {
	var sb strings.Builder
	for _, c := range r {
		sb.WriteRune(c.sym)
	}
	return sb.String()
}

type Grid []Row

func (g Grid) WithinBounds(x, y int) bool {
	if y >= len(g) {
		return false
	}
	if x >= len(g[y]) {
		return false
	}
	return true
}

func (g Grid) String() string {
	var sb strings.Builder
	for _, r := range g {
		sb.WriteString(r.String())
		sb.WriteRune('\n')
	}
	return sb.String()
}

func readGridFromReader(r io.Reader) (grid Grid, err error) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		var row Row
		for _, c := range line {
			row = append(row, Cell{sym: c})
		}
		for n := 0; n < 10; n++ {
			row = append(row, row...)
		}
		grid = append(grid, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading standard input: %w", err)
	}

	return
}

func countTreesInPath(grid Grid) int {
	var trees int

	x, y := 0, 0

	for grid.WithinBounds(x, y) {
		if grid[y][x].Tree() {
			trees++
		}
		x, y = x+3, y+1
	}

	return trees
}

func main() {
	grid, err := readGridFromReader(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading input: %s", err)
		os.Exit(2)
	}

	trees := countTreesInPath(grid)
	fmt.Printf("%d\n", trees)
}
