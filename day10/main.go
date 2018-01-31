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

func part2(input string) string {
	lengths := []int{}
	for l := range input {
		lengths = append(lengths, int(input[l]))
	}
	lengths = append(lengths, []int{17, 31, 73, 47, 23}...)

	list := []int{}
	for i := 0; i < 256; i++ {
		list = append(list, i)
	}

	pos := 0
	skip := 0
	for round := 0; round < 64; round++ {
		for _, l := range lengths {

			reverse := make([]int, l)
			for k := 0; k < l; k++ {
				reverse[k] = list[(pos+k)%256]
			}
			for k := 0; k < l; k++ {
				list[(pos+k)%256] = reverse[l-k-1]
			}
			pos = (pos + l + skip) % 256
			skip++
		}
	}

	var hash [16]byte
	for c := 0; c < 16; c++ {
		h := byte(list[c*16+0])
		for v := 1; v < 16; v++ {
			h = h ^ byte(list[c*16+v])
		}
		hash[c] = h
	}

	return fmt.Sprintf("%x", hash)
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	data := satoia(strings.Split(string(input), ","))
	fmt.Printf("Part1: %d\n", part1(256, data))
	fmt.Printf("Part2: %s\n", part2(strings.Trim(string(input), "\n")))
}
