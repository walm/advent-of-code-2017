package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// satoia string array to int array ;)
func satoia(s []string) []int {
	ns := []int{}
	for _, c := range s {
		if n, err := strconv.Atoi(c); err == nil {
			ns = append(ns, n)
		}
	}
	return ns
}

// minMax values from a int array
func minMax(a []int) (min, max int) {
	for _, n := range a {
		if n <= min || min == 0 {
			min = n
		}
		if n >= max {
			max = n
		}
	}
	return
}

func part1(sheet [][]int) (n int) {
	for _, r := range sheet {
		min, max := minMax(r)
		n += max - min
	}
	return
}

func main() {
	sheet := [][]int{}
	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")
	for _, col := range rows {
		cols := satoia(strings.Split(col, "\t"))
		sheet = append(sheet, cols)
	}
	fmt.Printf("Part1 sum: %d\n", part1(sheet))
}
