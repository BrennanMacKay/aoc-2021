package d5

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

type line struct {
	start point
	end   point
}

var regx = regexp.MustCompile(`(\d*),(\d*) -> (\d*),(\d*)`)

func Day5(args []string) int {
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
	lines := parseLines(input)

	var m = make(map[point]int, 0)

	for _, l := range lines {
		points := lineToPoints(l, false)
		for _, p := range points {
			m[p]++
		}
	}

	fmt.Println(m)
	count := 0
	for k, v := range m {
		if v > 1 {
			fmt.Println("DOUBLE", v, k)
			count++
		}
	}

	fmt.Println(count)
	return 0
}

func part2(args []string) int {
	input := tools.Read(args[0])
	lines := parseLines(input)

	var m = make(map[point]int, 0)

	for _, l := range lines {
		points := lineToPoints(l, true)
		for _, p := range points {
			m[p]++
		}
	}

	fmt.Println(m)
	count := 0
	for k, v := range m {
		if v > 1 {
			fmt.Println("DOUBLE", v, k)
			count++
		}
	}

	fmt.Println(count)
	return 0
}

func lineToPoints(l line, diag bool) []point {
	var points = make([]point, 0)

	// vertical
	if l.start.x == l.end.x {
		var dir int
		if l.start.y < l.end.y {
			dir = 1
		} else {
			dir = -1
		}

		for i := l.start.y; i != l.end.y+dir; i += dir {
			points = append(points, point{l.start.x, i})
		}
		fmt.Println("V POINTS ON LINE", l.start, l.end, points)

	} else if l.start.y == l.end.y {
		// horizontal
		var dir int
		if l.start.x < l.end.x {
			dir = 1
		} else {
			dir = -1
		}

		for i := l.start.x; i != l.end.x+dir; i += dir {
			points = append(points, point{i, l.start.y})
		}
		fmt.Println("H POINTS ON LINE", l.start, l.end, points)

	} else if diag {
		// diag always has a slope of 1
		var xDir, yDir int
		if l.start.x < l.end.x {
			xDir = 1
		} else {
			xDir = -1
		}

		if l.start.y < l.end.y {
			yDir = 1
		} else {
			yDir = -1
		}

		xIdx, yIdx := l.start.x, l.start.y
		for xIdx != l.end.x+xDir && yIdx != l.end.y+yDir {
			points = append(points, point{xIdx, yIdx})
			xIdx += xDir
			yIdx += yDir
		}

		fmt.Println("D POINTS ON LINE", l.start, l.end, points)
	} else {
		fmt.Println("D POINTS ON LINE", l.start, l.end, points)
	}

	return points
}

func parseLines(input []string) []line {
	lines := make([]line, len(input))
	for i, line := range input {
		lines[i] = parseLine(line)
	}

	return lines
}

func parseLine(input string) line {
	sParts := regx.FindStringSubmatch(input)
	parts := make([]int, len(sParts)-1)
	for i, s := range sParts[1:] {
		parts[i], _ = strconv.Atoi(s)
	}

	start := point{parts[0], parts[1]}
	end := point{parts[2], parts[3]}

	fmt.Println("PARSED LINE", start, end, parts)
	return line{start, end}
}
