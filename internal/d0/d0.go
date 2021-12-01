package d0

import "fmt"

func Day0(args []string) int {
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
	fmt.Printf("Part 1: %v\n", args)
	return 0
}

func part2(args []string) int {
	fmt.Printf("Part 2: %v\n", args)
	return 0
}
