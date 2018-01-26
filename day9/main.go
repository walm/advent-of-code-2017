package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func part1(data string) int {
	score := 0
	level := 0
	garbage := false
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
		}

		if c == "<" {
			garbage = true
		}
		if c == ">" {
			garbage = false
		}
		if c == "!" {
			skip = true
		}
	}
	return score
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	data := string(input)
	fmt.Printf("Part1: %d\n", part1(data))
}
