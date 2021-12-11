package d11

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"strconv"
)

type point struct {
	x int
	y int
}

var grid [10][10]int

func Day11(args []string) int {
	switch args[0] {
	case "p1":
		return part1(args[1:])
	case "p2":
		return part2(args[1:])
	default:
		fmt.Println("Unknown part", args[0])
		return 1
	}
}

func part1(args []string) int {
	input := tools.Read(args[0])
	parseGrid(input)

	dur, _ := strconv.Atoi(args[1])
	flashes := 0
	for i := 0; i < dur; i++ {
		flashesInStep := step()
		flashes += flashesInStep
		fmt.Printf("Step %d Flashes %d Total %d\n", i+1, flashesInStep, flashes)
		printGrid()
	}

	return 0
}

func part2(args []string) int {
	input := tools.Read(args[0])
	parseGrid(input)

	flashes := 0
	for i := 0; ; i++ {
		flashesInStep := step()

		flashes += flashesInStep
		fmt.Printf("Step %d Flashes %d Total %d\n", i+1, flashesInStep, flashes)
		printGrid()

		if flashesInStep == 100 {
			break
		}

	}

	return 0
}

func step() int {
	flashes := 0

	// initial increment
	for y, row := range grid {
		for x := range row {
			grid[y][x]++
		}
	}

	// steps
	for y, row := range grid {
		for x, energy := range row {
			if energy > 9 {
				flashes += stepOn(point{x, y})
			}
		}
	}

	return flashes
}

func stepOn(p point) int {
	energy := grid[p.y][p.x]
	if energy >= 9 { // we're looking for greater than, but we increment
		grid[p.y][p.x] = 0
		adjs := adjacent(p)
		acc := 1
		for _, adj := range adjs {
			acc += stepOn(adj)
		}

		return acc
	} else if energy != 0 { // we reset to 0 immediately, don't increment if zero
		grid[p.y][p.x]++

		return 0
	} else {
		return 0
	}
}

func adjacent(p point) []point {
	points := make([]point, 0)
	if p.y > 0 {
		points = append(points, point{p.x, p.y - 1})
		if p.x > 0 {
			points = append(points, point{p.x - 1, p.y - 1})
		}
		if p.x < len(grid)-1 {
			points = append(points, point{p.x + 1, p.y - 1})
		}
	}

	if p.x > 0 {
		points = append(points, point{p.x - 1, p.y})
		if p.y < len(grid)-1 {
			points = append(points, point{p.x - 1, p.y + 1})
		}
	}

	if p.y < len(grid)-1 {
		points = append(points, point{p.x, p.y + 1})
		if p.x < len(grid)-1 {
			points = append(points, point{p.x + 1, p.y + 1})
		}
	}

	if p.x < len(grid)-1 {
		points = append(points, point{p.x + 1, p.y})
	}

	return points
}

func parseGrid(input []string) {
	for y, row := range input {
		for x, en := range []rune(row) {
			grid[y][x] = int(en - '0')
		}
	}
}

func printInput(input []string) {
	for _, row := range input {
		fmt.Println(row)
	}
}

func printGrid() {
	for y := range grid {
		for x := range grid[y] {
			fmt.Print(grid[y][x])
		}
		fmt.Println()
	}
}
