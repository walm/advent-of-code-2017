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
		if n, err := strconv.Atoi(strings.Trim(c, "")); err == nil {
			ns = append(ns, n)
		}
	}
	return ns
}

func indexAndMax(a []int) (i int, m int) {
	for p, n := range a {
		if n > m {
			m = n
			i = p
		}
	}
	return
}

func sjoin(a []int) (s string) {
	s = fmt.Sprintf("%v", a)
	return
}

func redist(a []int, i int) []int {
	l := len(a)
	if i >= l {
		return a
	}
	n := a[i]
	if n == 0 {
		return a
	}
	a[i] = 0
	for {
		if i < l-1 {
			i++
		} else {
			i = 0
		}
		a[i] += 1
		n--
		if n <= 0 {
			break
		}
	}
	return a
}

func part1(blocks []int) (j int) {
	done := map[string]bool{
		sjoin(blocks): true,
	}
	for {
		j++
		i, _ := indexAndMax(blocks)
		blocks = redist(blocks, i)
		bs := sjoin(blocks)
		if done[bs] {
			return
		}
		done[bs] = true
	}
	return
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	blocks := satoia(strings.Split(strings.Trim(string(input), "\n"), "\t"))
	fmt.Printf("Part1: %d\n", part1(blocks))
}
