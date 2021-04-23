package two_sum

import (
	"fmt"
	"reflect"
	"testing"
)

type para struct {
	nums   []int
	target int
}

type prob1 struct {
	input  para
	output []int
}

var test = []prob1{
	{
		input: para{
			nums:   []int{2, 7, 11, 15},
			target: 9,
		},
		output: []int{0, 1},
	},
	{
		input: para{
			nums:   []int{3, 2, 4},
			target: 6,
		},
		output: []int{1, 2},
	},
	{
		input: para{
			nums:   []int{3, 3},
			target: 6,
		},
		output: []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	fmt.Printf("----------------Leetcode Problem 1------------------\n")

	for _, p := range test {
		in, expected := p.input, p.output
		actual := twoSum(in.nums, in.target)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("twoSum(%v, %d) = %v; expected %v", in.nums, in.target, actual, expected)
		}

	}
}

func TestTwoSumUsingMap(t *testing.T) {
	fmt.Printf("----------------Leetcode Problem 1------------------\n")

	for _, p := range test {
		in, expected := p.input, p.output
		actual := twoSumUsingMap(in.nums, in.target)
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("twoSumUsingMap(%v, %d) = %v; expected %v", in.nums, in.target, actual, expected)
		}
	}
}
