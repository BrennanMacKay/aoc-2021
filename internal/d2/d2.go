package d2

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"strconv"
	"strings"
)

type instruction struct {
	dir  string
	dist int
}

func Day2(args []string) int {
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
	var instructions = toInstructions(tools.Read(args[0]))
	var position = 0
	var depth = 0

	for _, v := range instructions {
		switch v.dir {
		case "forward":
			position += v.dist
		case "down":
			depth += v.dist
		case "up":
			depth -= v.dist
		}
	}

	fmt.Printf("Position: %d Depth: %d\n", position, depth)
	fmt.Printf("Answer: %d\n", position*depth)

	return 0
}

func part2(args []string) int {
	var instructions = toInstructions(tools.Read(args[0]))
	var position = 0
	var depth = 0
	var aim = 0

	for _, v := range instructions {
		switch v.dir {
		case "forward":
			position += v.dist
			depth += aim * v.dist
		case "down":
			aim += v.dist
		case "up":
			aim -= v.dist
		}
	}

	fmt.Printf("Position: %d Depth: %d\n", position, depth)
	fmt.Printf("Answer: %d\n", position*depth)

	return 0
}

func toInstructions(input []string) []instruction {
	instructions := make([]instruction, len(input))

	for i, v := range input {
		parts := strings.Fields(v)
		var dist, err = strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		instructions[i] = instruction{parts[0], dist}
	}

	return instructions
}
