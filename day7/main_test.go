package main

import "testing"

var testRows = []string{
	"pbga (66)",
	"xhth (57)",
	"ebii (61)",
	"havc (66)",
	"ktlj (57)",
	"fwft (72) -> ktlj, cntj, xhth",
	"qoyq (66)",
	"padx (45) -> pbga, havc, qoyq",
	"tknk (41) -> ugml, padx, fwft",
	"jptl (61)",
	"ugml (68) -> gyxo, ebii, jptl",
	"gyxo (61)",
	"cntj (57)",
}

func Test_parseRow(t *testing.T) {
	n, w, rel := parseRow("ugml (68) -> gyxo, ebii, jptl")
	expn := "ugml"
	expw := 68
	if n != expn || w != expw || len(rel) != 3 || rel[0] != "gyxo" {
		t.Errorf("result missmatch, got: %s %d %v, expect: %s %d.", n, w, rel, expn, expw)
	}

	n, w, rel = parseRow("bmke (8)")
	expn = "bmke"
	expw = 8
	if n != expn || w != expw || len(rel) != 0 {
		t.Errorf("result missmatch, got: %s %d, expect: %s %d.", n, w, expn, expw)
	}
}

func TestPart1_1(t *testing.T) {
	res := part1(testRows)
	exp := "tknk"
	if res != exp {
		t.Errorf("result missmatch, got: %s, expect: %s.", res, exp)
	}
}

func TestPart2_1(t *testing.T) {
	res := part2(testRows)
	exp := 60
	if res != exp {
		t.Errorf("result missmatch, got: %d, expect: %d.", res, exp)
	}
}
