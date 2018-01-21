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

func part1(maze []int) (j int) {
	pos := 0
	ml := len(maze)
	for {
		n := maze[pos]
		maze[pos] = n + 1
		pos += n
		j++
		if pos >= ml || pos < 0 {
			break
		}
	}
	return
}

func part2(maze []int) (j int) {
	pos := 0
	ml := len(maze)
	for {
		n := maze[pos]
		if n >= 3 {
			maze[pos] = n - 1
		} else {
			maze[pos] = n + 1
		}
		pos += n
		j++
		// fmt.Printf("m:%v pos:%d j:%d\n", maze, pos, j)
		if pos >= ml || pos < 0 {
			break
		}
	}
	return
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	maze := satoia(strings.Split(string(input), "\n"))
	fmt.Printf("Part1: %d\n", part1(maze))

	// recreate the maze as it change in part1 ;)
	maze = satoia(strings.Split(string(input), "\n"))
	fmt.Printf("Part2: %d\n", part2(maze))
}
