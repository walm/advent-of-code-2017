package main

import "testing"

func TestPart1(t *testing.T) {
	tests := map[string]bool{
		"aa bb cc dd ee":  true,
		"aa bb cc dd aa":  false,
		"aa bb cc dd aaa": true,
	}

	for p, exp := range tests {
		res := valid(p)
		if res != exp {
			t.Errorf("pass '%s' got: %v, expect: %v.", p, res, exp)
		}
	}
}

func TestPart1_1(t *testing.T) {
	ps := []string{
		"aa bb cc dd ee",
		"aa bb cc dd aa",
		"aa bb cc dd aaa",
	}
	res := part1(ps)
	exp := 2
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}
