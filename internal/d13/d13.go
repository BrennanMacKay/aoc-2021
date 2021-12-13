package d13

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"regexp"
	"strconv"
)

type point struct {
	x int
	y int
}

type fold struct {
	axis  string
	value int
}

var pRegex = regexp.MustCompile(`(\d*),(\d*)`)
var fRegex = regexp.MustCompile(`fold along (.)=(\d*)`)

type grid [][]bool

func Day13(args []string) int {
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
	points, folds := readInput(input)

	fmt.Println(points)
	fmt.Println(folds)

	bound := getBound(points)
	g := buildGrid(points, bound)
	printGrid(g)

	for i, f := range folds[0:1] {
		fmt.Println("Fold", i, f)
		switch f.axis {
		case "x":
			points = foldLeft(points, f.value)
			bound.x = f.value - 1
		case "y":
			points = foldUp(points, f.value)
			bound.y = f.value - 1
		}
		fmt.Println(bound)
		fmt.Println(points)

		g = buildGrid(points, bound)
		printGrid(g)

		points = mergePoints(points)
		fmt.Println(len(points))
	}

	return 0
}

func part2(args []string) int {
	input := tools.Read(args[0])
	points, folds := readInput(input)

	fmt.Println(points)
	fmt.Println(folds)

	bound := getBound(points)
	g := buildGrid(points, bound)
	printGrid(g)

	for i, f := range folds {
		fmt.Println("Fold", i, f)
		switch f.axis {
		case "x":
			points = foldLeft(points, f.value)
			bound.x = f.value - 1
		case "y":
			points = foldUp(points, f.value)
			bound.y = f.value - 1
		}
		fmt.Println(bound)
		fmt.Println(points)

		g = buildGrid(points, bound)
		printGrid(g)

		points = mergePoints(points)
		fmt.Println(len(points))
	}

	return 0
}

//This is kind of wasteful, but I don't want to switch to using a map everywhere
func mergePoints(points []point) []point {
	pMap := make(map[point]struct{}, len(points))

	for _, p := range points {
		pMap[p] = struct{}{}
	}

	newPoints := make([]point, len(pMap))

	i := 0
	for k := range pMap {
		newPoints[i] = k
		i++
	}

	return newPoints
}

func foldUp(points []point, fold int) []point {
	newPoints := make([]point, len(points))

	for i, p := range points {
		if p.y < fold {
			newPoints[i] = p
		} else {
			dist := p.y - fold
			newPoints[i] = point{p.x, fold - dist}
		}
	}

	return newPoints
}

func foldLeft(points []point, fold int) []point {
	newPoints := make([]point, len(points))

	for i, p := range points {
		if p.x < fold {
			newPoints[i] = p
		} else {
			dist := p.x - fold
			newPoints[i] = point{fold - dist, p.y}
		}
	}

	return newPoints

}

func getBound(points []point) point {
	maxX, maxY := 0, 0

	for _, p := range points {
		if p.x > maxX {
			maxX = p.x
		}

		if p.y > maxY {
			maxY = p.y
		}
	}

	return point{maxX, maxY}
}

func buildGrid(points []point, bound point) grid {
	g := make(grid, bound.y+1)
	for i := range g {
		g[i] = make([]bool, bound.x+1)
	}

	for _, p := range points {
		g[p.y][p.x] = true
	}

	return g
}

func readInput(input []string) ([]point, []fold) {
	points := make([]point, 0)
	folds := make([]fold, 0)

	for i, line := range input {
		switch {
		case pRegex.MatchString(line):
			parts := pRegex.FindStringSubmatch(line)
			x, _ := strconv.Atoi(parts[1])
			y, _ := strconv.Atoi(parts[2])
			points = append(points, point{x, y})
		case fRegex.MatchString(line):
			parts := fRegex.FindStringSubmatch(line)
			pos, _ := strconv.Atoi(parts[2])
			folds = append(folds, fold{parts[1], pos})
		default:
			fmt.Println("Unknown", line, i)
		}
	}

	return points, folds
}

func printGrid(g grid) {
	fmt.Println()
	for _, row := range g {
		for _, v := range row {
			if v {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
