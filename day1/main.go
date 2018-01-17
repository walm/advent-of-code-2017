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

func check(digits string) (sum int) {
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

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	digits := strings.Trim(string(input), "\n")
	fmt.Printf("Sum: %d", check(digits))
}
