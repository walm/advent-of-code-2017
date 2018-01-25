package main

import "testing"

var testRows = []string{
	"b inc 5 if a > 1",
	"a inc 1 if b < 5",
	"c dec -10 if a >= 1",
	"c inc -20 if c == 10",
}

func Test_parseRow(t *testing.T) {
	n, f, v, lv, o, rv := parseRow("c dec -10 if a >= 1")
	expn := "c"
	expf := "dec"
	expv := -10
	explv := "a"
	expo := ">="
	exprv := 1
	if n != expn || f != expf || v != expv || lv != explv || o != expo || rv != exprv {
		t.Errorf("result missmatch, got: %s %s %d %s %s %d, expect: %s %s %d %s %s %d.", n, f, v, lv, o, rv, expn, expf, expv, explv, expo, exprv)
	}
}

func TestPart1_1(t *testing.T) {
	res := part1(testRows)
	exp := 1
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}
