package main

import (
	"fmt"
	"math"
)

func part1(in int) int {
	x := 0
	y := 0
	v := 0
	sl := 1.0
	dir := 0
	for v < in {

		v++
		if v == in {
			break
		}

		switch dir {
		case 0: // down
			y++
		case 1: // right
			x++
		case 2: // up
			y--
		case 3: // left
			x--
		}

		if (dir%2 == 1 && math.Abs(float64(x)) == math.Ceil(sl/2)) ||
			(dir%2 == 0 && math.Abs(float64(y)) == math.Ceil(sl/2)) {
			dir++
			if dir == 4 {
				dir = 0
				sl += 2
			}
		}

	}
	return int(math.Abs(float64(x)) + math.Abs(float64(y)))
}

type pos struct {
	x int
	y int
}

func part2(in int) int {
	x := 0
	y := 0
	v := 1
	sl := 1.0
	dir := 0
	sqr := map[pos]int{}
	for v < in {

		sqr[pos{x, y}] = v

		if v >= in {
			break
		}

		switch dir {
		case 0: // down
			y++
		case 1: // right
			x++
		case 2: // up
			y--
		case 3: // left
			x--
		}

		v = sumNeighbors(sqr, x, y)

		if (dir%2 == 1 && math.Abs(float64(x)) == math.Ceil(sl/2)) ||
			(dir%2 == 0 && math.Abs(float64(y)) == math.Ceil(sl/2)) {
			dir++
			if dir == 4 {
				dir = 0
				sl += 2
			}
		}

	}
	return v
}

func sumNeighbors(sqr map[pos]int, x, y int) (out int) {
	for xl := x - 1; xl <= x+1; xl++ {
		for yl := y - 1; yl <= y+1; yl++ {
			out += sqr[pos{xl, yl}]
		}
	}
	return
}

func main() {
	input := 277678
	fmt.Printf("Part1: %d\n", part1(input))
	fmt.Printf("Part2: %d\n", part2(input))
}
