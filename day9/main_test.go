package main

import "testing"

func TestPart1_1(t *testing.T) {
	res := part1("{}")
	exp := 1
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_2(t *testing.T) {
	res := part1("{{{}}}")
	exp := 6
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_3(t *testing.T) {
	res := part1("{{},{}}")
	exp := 5
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_4(t *testing.T) {
	res := part1("{{{},{},{{}}}}")
	exp := 16
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_5(t *testing.T) {
	res := part1("{<a>,<a>,<a>,<a>}")
	exp := 1
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_6(t *testing.T) {
	res := part1("{{<ab>},{<ab>},{<ab>},{<ab>}}")
	exp := 9
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_7(t *testing.T) {
	res := part1("{{<!!>},{<!!>},{<!!>},{<!!>}}")
	exp := 9
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart1_8(t *testing.T) {
	res := part1("{{<a!>},{<a!>},{<a!>},{<ab>}}")
	exp := 3
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestPart2(t *testing.T) {
	tests := map[string]int{
		"<>": 0,
		"<random characters>": 17,
		"<<<<>":               3,
		"<{!>}>":              2,
		"<!!>":                0,
		"<!!!>>":              0,
		`<{o"i!a,<{i<a>`:      10,
	}
	for d, exp := range tests {
		res := part2(d)
		if res != exp {
			t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
		}
	}
}
