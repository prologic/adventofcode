package aoc

import (
	"fmt"
)

const (
	MaxRows = 128
	MaxCols = 8
)

type Boardingpass struct {
	Row int
	Col int
}

func (bp Boardingpass) SeatID() int {
	return (bp.Row * 8) + bp.Col
}

func (bp Boardingpass) String() string {
	return fmt.Sprintf("%dx%d:%d", bp.Row, bp.Col, bp.SeatID())
}

func ParseBoardingpass(s string) (bp Boardingpass, err error) {
	rows := []int{0, MaxRows}
	cols := []int{0, MaxCols}

	for _, c := range s[:7] {
		switch c {
		case 'F':
			rows[1] -= (rows[1] - rows[0]) / 2
		case 'B':
			rows[0] += (rows[1] - rows[0]) / 2
		}
	}

	if s[6] == 'F' {
		bp.Row = rows[0]
	} else {
		bp.Row = MinOf((MaxRows - 1), rows[1])
	}

	for _, c := range s[7:] {
		switch c {
		case 'L':
			cols[1] -= (cols[1] - cols[0]) / 2
		case 'R':
			cols[0] += (cols[1] - cols[0]) / 2
		}
	}

	if s[9] == 'L' {
		bp.Col = cols[0]
	} else {
		bp.Col = MinOf((MaxCols - 1), cols[1])
	}

	return
}
