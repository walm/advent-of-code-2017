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

func part1(size int, lengths []int) int {
	list := []int{}
	for i := 0; i < size; i++ {
		list = append(list, i)
	}
	ll := len(list)

	start := 0
	skip := 0
	wl := map[int]int{}
	for _, l := range lengths {
		end := start + l - 1
		for i := 0; i < l; i++ {
			pos := start + i
			if pos >= ll {
				pos = pos - ll
			}
			npos := end - i
			if npos >= ll {
				npos = npos - ll
			}
			wl[npos] = list[pos]
		}

		for i, n := range wl {
			list[i] = n
		}
		start = start + l + skip
		if start >= ll {
			start = start - ll
		}
		skip++
	}
	return list[0] * list[1]
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	data := satoia(strings.Split(string(input), ","))
	fmt.Printf("Part1: %d\n", part1(256, data))
}
