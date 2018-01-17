package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func btoi(b byte) int {
	n, _ := strconv.Atoi(string(b))
	return n
}

func part1(digits string) (sum int) {
	for i, _ := range digits {
		if i > 0 && digits[i] == digits[i-1] {
			sum += btoi(digits[i])
		}
	}
	if digits[0] == digits[len(digits)-1] {
		sum += btoi(digits[0])
	}
	return
}

func part2(digits string) (sum int) {
	steps := len(digits) / 2
	for i, _ := range digits {
		nextPos := i + steps
		if nextPos >= len(digits) {
			nextPos = nextPos - len(digits)
		}
		if digits[i] == digits[nextPos] {
			sum += btoi(digits[i])
		}
	}
	return
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	digits := strings.Trim(string(input), "\n")
	fmt.Printf("Part1 sum: %d\n", part1(digits))
	fmt.Printf("Part2 sum: %d\n", part2(digits))
}
