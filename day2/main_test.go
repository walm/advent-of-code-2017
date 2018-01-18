package main

import "testing"

var testSheet = [][]int{
	0: []int{5, 1, 9, 5},
	1: []int{7, 5, 3},
	2: []int{2, 4, 6, 8},
}

func TestStoi(t *testing.T) {
	s := []string{"1", "2", "3"}
	res := satoia(s)
	if len(res) != len(s) {
		t.Errorf("result missmatch, length should be %d", len(s))
	}
	if res[0] != 1 {
		t.Errorf("result missmatch, first value got: %d expected: 1", res[0])
	}
}

func TestPart1(t *testing.T) {
	res := part1(testSheet)
	exp := 18
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}
