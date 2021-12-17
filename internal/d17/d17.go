package d17

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"math"
	"regexp"
	"strconv"
)

var targetRegex = regexp.MustCompile(`target area: x=(-?\d*)..(-?\d*), y=(-?\d*)..(-?\d*)`)

type point struct {
	x int
	y int
}

type target struct {
	topLeft     point
	bottomRight point
}

type projectile struct {
	xV int
	yV int

	maxHeight int
	position  point
}

func Day17(args []string) int {
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

	target := readTarget(tools.Read(args[0])[0])
	minX := minX(target.topLeft.x)
	maxY := maxY(target.bottomRight.y)
	fmt.Println(minX, maxY)
	projectile := projectile{minX, maxY, 0, point{0, 0}}
	max, hit, pos := step(projectile, target, make([]point, 0))
	plot(pos, target)
	fmt.Println(max, hit)

	return 0
}

func part2(args []string) int {
	//points := make([]point, 0)
	count := 0
	target := readTarget(tools.Read(args[0])[0])
	for y := -1000; y <= 1000; y++ {
		if y%10 == 0 {
			fmt.Println(y)
		}
		for x := -1000; x <= 1000; x++ {
			projectile := projectile{x, y, 0, point{0, 0}}
			if _, hit, _ := step(projectile, target, make([]point, 0)); hit {
				//points = append(points, point{x, y})
				count++
			}
		}
	}

	fmt.Println("Result", count)

	return 0
}

func minX(tX int) int {
	x := 0

	for tX >= 0 {
		tX -= x
		x++
	}

	return x - 1
}

func maxX(tX int) int {
	x := 0

	for tX >= 0 {
		tX -= x
		x++
	}

	fmt.Println(x)
	return x - 2
}

func maxY(tY int) int {
	return (tY * -1) - 1
}

func step(proj projectile, target target, positions []point) (int, bool, []point) {
	newProj := projectile{}
	newProj.position = point{proj.position.x + proj.xV, proj.position.y + proj.yV}
	newProj.xV, newProj.yV = newXV(proj.xV), newYV(proj.yV)
	newProj.maxHeight = proj.maxHeight
	if newProj.position.y > proj.maxHeight {
		newProj.maxHeight = newProj.position.y
	}

	positions = append(positions, newProj.position)
	if inTarget(newProj.position, target) {
		return newProj.maxHeight, true, positions
	}

	prevDist := distance(proj.position, target)
	newDist := distance(newProj.position, target)
	if newDist > prevDist && newProj.yV < 0 {
		return 0, false, positions
	} else {
		return step(newProj, target, positions)
	}
}

func newXV(x int) int {
	if x == 0 {
		return 0
	} else if x > 0 {
		return x - 1
	} else {
		return x + 1
	}
}

func newYV(y int) int {
	return y - 1
}

func distance(point point, target target) int {
	a := point.x - target.bottomRight.x
	b := point.y - target.bottomRight.y

	c := math.Sqrt(float64(a*a + b*b))

	// not too concerned about precision
	return int(c)
}

func inTarget(point point, target target) bool {
	if target.topLeft.x <= point.x && point.x <= target.bottomRight.x {
		if target.topLeft.y >= point.y && point.y >= target.bottomRight.y {
			return true
		}
	}

	return false
}

func readTarget(input string) target {
	parts := targetRegex.FindStringSubmatch(input)
	xMin, _ := strconv.Atoi(parts[1])
	xMax, _ := strconv.Atoi(parts[2])
	yMin, _ := strconv.Atoi(parts[3])
	yMax, _ := strconv.Atoi(parts[4])

	return target{point{xMin, yMax}, point{xMax, yMin}}
}

func plot(positions []point, target target) {
	maxX := 0
	maxY := 0
	minY := math.MaxInt

	for _, pos := range positions {
		if pos.x > maxX {
			maxX = pos.x
		}
		if pos.y > maxY {
			maxY = pos.y
		}
		if pos.y < minY {
			minY = pos.y
		}
	}

	if target.bottomRight.y < minY {
		minY = target.bottomRight.y
	}

	if target.bottomRight.x > maxX {
		maxX = target.bottomRight.x
	}

	xDist := maxX
	yDist := maxY - minY

	grid := make([][]rune, yDist+1)
	for i := range grid {
		grid[i] = make([]rune, xDist+1)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			grid[y][x] = '.'
		}
	}

	yDiff := minY * -1
	for y := target.topLeft.y + yDiff; y >= target.bottomRight.y+yDiff; y-- {
		for x := target.topLeft.x; x <= target.bottomRight.x; x++ {
			grid[y][x] = 'T'
		}
	}

	for _, p := range positions {
		grid[p.y+yDiff][p.x] = '#'
	}

	grid[yDiff][0] = 'S'

	for y := len(grid) - 1; y >= 0; y-- {
		fmt.Println(string(grid[y]))
	}
}
