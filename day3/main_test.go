package main

import "testing"

func TestPart1_1(t *testing.T) {
	res := part1(1)
	exp := 0
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_2(t *testing.T) {
	res := part1(12)
	exp := 3
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_3(t *testing.T) {
	res := part1(23)
	exp := 2
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_4(t *testing.T) {
	res := part1(1024)
	exp := 31
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}
