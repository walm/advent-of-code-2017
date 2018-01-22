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

func Test_indexAndMax(t *testing.T) {
	b := []int{0, 2, 7, 0}
	i, m := indexAndMax(b)
	if i != 2 && m != 7 {
		t.Errorf("result missmatch, got: i:%d m:%d, expect: i:%d m:%d.", i, m, 2, 7)
	}
}

func Test_sjoin(t *testing.T) {
	b := []int{0, 2, 7, 0}
	res := sjoin(b)
	exp := "[0 2 7 0]"
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func Test_redist(t *testing.T) {
	b := []int{0, 2, 7, 0}
	res := redist(b, 2)
	exp := []int{2, 4, 1, 2}
	if sjoin(res) != sjoin(exp) {
		t.Errorf("result missmatch, got: %v, expect: %v.", res, exp)
	}
}

func TestPart1_1(t *testing.T) {
	b := []int{0, 2, 7, 0}
	res := part1(b)
	exp := 5
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart2_1(t *testing.T) {
	b := []int{0, 2, 7, 0}
	res := part2(b)
	exp := 4
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}
