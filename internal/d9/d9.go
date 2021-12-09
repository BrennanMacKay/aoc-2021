package d9

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"sort"
)

type point struct {
	x int
	y int
}

var grid [][]int

func Day9(args []string) int {
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
	printInput(input)
	parseGrid(input)
	printGrid()
	lows := locateLows()
	fmt.Println(lows)
	fmt.Println(riskLevel(lows))
	return 0
}

func part2(args []string) int {
	input := tools.Read(args[0])
	printInput(input)
	parseGrid(input)
	printGrid()
	lows := locateLows()
	fmt.Println(lows)

	basinSizes := make([]int, len(lows))
	for i, low := range lows {
		basinSizes[i] = basinSize(low)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(basinSizes)))
	fmt.Println(basinSizes)

	fmt.Println(basinSizes[0] * basinSizes[1] * basinSizes[2])
	return 0
}

func basinSize(low point) int {
	visited := make(map[point]struct{}, 0)

	return visit(visited, low)
}

func visit(visited map[point]struct{}, p point) int {
	if grid[p.y][p.x] == 9 {
		return 0
	}

	visited[p] = struct{}{}
	adj := [4]point{{p.x, p.y - 1}, {p.x, p.y + 1}, {p.x - 1, p.y}, {p.x + 1, p.y}}

	sum := 1
	for _, adjP := range adj {
		if _, ok := visited[adjP]; !ok {
			sum += visit(visited, adjP)
		}
	}

	return sum
}

func locateLows() []point {
	lows := make([]point, 0)
	for y := 1; y < len(grid)-1; y++ {
		for x := 1; x < len(grid[y])-1; x++ {
			var adj = [4]int{grid[y-1][x], grid[y+1][x], grid[y][x-1], grid[y][x+1]}
			lowest := true
			for _, other := range adj {
				if other <= grid[y][x] {
					lowest = false
					break
				}
			}

			if lowest {
				lows = append(lows, point{x, y})
			}
		}
	}

	return lows
}

func riskLevel(lows []point) int {
	sum := 0
	for _, low := range lows {
		sum += grid[low.y][low.x] + 1
	}

	return sum
}

func parseGrid(input []string) {
	rows := len(input)
	columns := len(input[0])

	// Pad edges to simplify logic
	grid = make([][]int, rows+2)
	for y := range grid {
		grid[y] = make([]int, columns+2)
		grid[y][0] = 9
		grid[y][columns+1] = 9
		if y == 0 || y == rows+1 {
			for x := range grid[y] {
				grid[y][x] = 9
			}
		}
	}

	for y, row := range input {
		for x, height := range []rune(row) {
			grid[y+1][x+1] = int(height - '0')
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
