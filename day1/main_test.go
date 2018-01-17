package main

import "testing"

func TestCheck1(t *testing.T) {
	res := check("1122")
	exp := 3
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestCheck2(t *testing.T) {
	res := check("1111")
	exp := 4
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestCheck3(t *testing.T) {
	res := check("1234")
	exp := 0
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}

func TestCheck4(t *testing.T) {
	res := check("91212129")
	exp := 9
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}
