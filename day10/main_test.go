package main

import "testing"

func TestPart1_1(t *testing.T) {
	res := part1(5, []int{3, 4, 1, 5})
	exp := 12
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}
