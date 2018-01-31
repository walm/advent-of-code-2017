package main

import "testing"

func TestPart1_1(t *testing.T) {
	res := part1(5, []int{3, 4, 1, 5})
	exp := 12
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart2_1(t *testing.T) {
	tests := map[string]string{
		"":         "a2582a3a0e66e6e86e3812dcb672a272",
		"AoC 2017": "33efeb34ea91902bb2f59c9920caa6cd",
		"1,2,3":    "3efbe78a8d82f29979031a4aa0b16a9d",
		"1,2,4":    "63960835bcdc130f0b66d7ff4f6a5a8e",
	}
	for inp, exp := range tests {
		res := part2(inp)
		if res != exp {
			t.Errorf("result missmatch, got: %s, expect: %s.", res, exp)
		}
	}
}
