package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func parseRow(s string) (n, f string, v int, lv, o string, rv int) {
	rx := regexp.MustCompile(`^(\w+)\s(\w+)\s(-?\d+)\s\w+\s(\w+)\s([<>=!]+)\s(-?\d+)$`)
	m := rx.FindAllStringSubmatch(s, -1)
	n = m[0][1]
	f = m[0][2]
	v, _ = strconv.Atoi(m[0][3])
	lv = m[0][4]
	o = m[0][5]
	rv, _ = strconv.Atoi(m[0][6])
	return
}

func parseVars(rows []string) (int, int) {
	vars := map[string]int{}
	max := 0
	for _, r := range rows {
		if r == "" {
			continue
		}
		n, f, v, lv, o, rv := parseRow(r)
		do := false
		if o == "==" && vars[lv] == rv {
			do = true
		}
		if o == "!=" && vars[lv] != rv {
			do = true
		}
		if o == ">" && vars[lv] > rv {
			do = true
		}
		if o == ">=" && vars[lv] >= rv {
			do = true
		}
		if o == "<" && vars[lv] < rv {
			do = true
		}
		if o == "<=" && vars[lv] <= rv {
			do = true
		}
		if do {
			if f == "inc" {
				vars[n] = vars[n] + v
			} else {
				vars[n] = vars[n] - v
			}
			if vars[n] > max {
				max = vars[n]
			}
		}
	}
	largest := 0
	for _, v := range vars {
		if v > largest {
			largest = v
		}
	}
	return largest, max
}

func part1(rows []string) int {
	l, _ := parseVars(rows)
	return l
}

func part2(rows []string) int {
	_, m := parseVars(rows)
	return m
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")
	l, m := parseVars(rows)
	fmt.Printf("Part1: %d\n", l)
	fmt.Printf("Par21: %d\n", m)
}
