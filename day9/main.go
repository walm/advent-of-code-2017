package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func parse(data string) (int, int) {
	score := 0
	level := 0
	garbage := false
	gcount := 0
	skip := false
	for _, c := range strings.Split(data, "") {

		if skip {
			skip = false
			continue
		}

		if !garbage {
			if c == "{" {
				level += 1
				score += level
			}
			if c == "}" {
				level -= 1
			}
			if c == "<" {
				garbage = true
				continue
			}
		}

		if c == ">" {
			garbage = false
			continue
		}
		if c == "!" {
			skip = true
			continue
		}

		if garbage && !skip {
			gcount += 1
		}
	}
	return score, gcount
}

func part1(data string) int {
	score, _ := parse(data)
	return score
}

func part2(data string) int {
	_, g := parse(data)
	return g
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	data := string(input)
	score, garbage := parse(data)
	fmt.Printf("Part1: %d\n", score)
	fmt.Printf("Part2: %d\n", garbage)
}
