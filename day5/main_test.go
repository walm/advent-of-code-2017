package main

import "testing"

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

func TestPart1_1(t *testing.T) {
	maze := []int{0, 3, 0, 1, -3}
	res := part1(maze)
	exp := 5
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}
