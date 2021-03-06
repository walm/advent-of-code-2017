package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

func valid(p string) bool {
	exists := map[string]bool{}
	words := strings.Split(p, " ")
	for _, w := range words {
		if exists[w] {
			return false
		}
		exists[w] = true
	}
	return true
}

func part1(ps []string) int {
	count := 0
	for _, p := range ps {
		if p != "" && valid(p) {
			count++
		}
	}
	return count
}

func valid2(p string) bool {
	exists := map[string]bool{}
	words := strings.Split(p, " ")
	for _, w := range words {
		if w != "" {
			sw := strings.Split(w, "")
			sort.Strings(sw)
			s := strings.Join(sw, "")
			if exists[s] {
				return false
			}
			exists[s] = true
		}
	}
	return true
}

func part2(ps []string) int {
	count := 0
	for _, p := range ps {
		if p != "" && valid2(p) {
			count++
		}
	}
	return count
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")
	fmt.Printf("Part1: %d\n", part1(rows))
	fmt.Printf("Part2: %d\n", part2(rows))
}
