package aoc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBoardingpass(t *testing.T) {
	assert := assert.New(t)

	testCases := []struct {
		input  string
		row    int
		col    int
		seatid int
	}{
		{
			"BFFFBBFRRR",
			70, 7, 567,
		}, {
			"FFFBBBFRRR",
			14, 7, 119,
		}, {
			"BBFFBBFRLL",
			102, 4, 820,
		},
	}

	for _, testCase := range testCases {
		bp, err := ParseBoardingpass(testCase.input)
		assert.NoError(err)

		assert.Equal(testCase.row, bp.Row, "Row mismatch")
		assert.Equal(testCase.col, bp.Col, "Col mismatch")
		assert.Equal(testCase.seatid, bp.SeatID(), "SeatID mismatch")
	}
}
