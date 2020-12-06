package aoc

import (
	"fmt"
	"strconv"
	"strings"
)

type Boardingpass int

func (bp Boardingpass) Row() int {
	return int(bp) >> 3
}

func (bp Boardingpass) Col() int {
	return int(bp) & 0x7
}

func (bp Boardingpass) SeatID() int {
	return int(bp)
}

func (bp Boardingpass) String() string {
	return fmt.Sprintf("%dx%d:%d", bp.Row(), bp.Col(), bp.SeatID())
}

func ParseBoardingpass(s string) (bp Boardingpass, err error) {
	r := strings.NewReplacer("F", "0", "B", "1", "L", "0", "R", "1")
	b := r.Replace(s)

	n, err := strconv.ParseInt(b, 2, 64)
	if err != nil {
		return Boardingpass(0), fmt.Errorf("error parsing boarding pass: %w", err)
	}

	id := int(n)

	bp = Boardingpass(id)

	return
}
