package problem0001

import (
	"testing"

	"gotest.tools/assert"
)

type input struct {
	set    []int
	target int
}

type result struct {
	res []int
}

type problem struct {
	in  input
	res result
}

func TestTwoSum(t *testing.T) {
	table := []problem{
		problem{
			in: input{
				set:    []int{2, 7, 11, 15},
				target: 9,
			},
			res: result{
				res: []int{0, 1},
			},
		},
		problem{
			in: input{
				set:    []int{2, 7, 11, 15},
				target: 30,
			},
			res: result{
				res: nil,
			},
		},
	}

	for _, tb := range table {
		in, res := tb.in, tb.res
		assert.Equal(t, res.res, TwoSum(in.set, in.target))
	}
}
