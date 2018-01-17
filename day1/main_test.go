package main

import "testing"

func TestPart1_1(t *testing.T) {
	res := part1("1122")
	exp := 3
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_2(t *testing.T) {
	res := part1("1111")
	exp := 4
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_3(t *testing.T) {
	res := part1("1234")
	exp := 0
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_4(t *testing.T) {
	res := part1("91212129")
	exp := 9
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart2_1(t *testing.T) {
	res := part2("1212")
	exp := 6
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart2_2(t *testing.T) {
	res := part2("1221")
	exp := 0
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart2_3(t *testing.T) {
	res := part2("123425")
	exp := 4
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart2_4(t *testing.T) {
	res := part2("123123")
	exp := 12
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart2_5(t *testing.T) {
	res := part2("12131415")
	exp := 4
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}
