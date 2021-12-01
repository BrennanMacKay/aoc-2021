package d1

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
)

func Day1(args []string) int {
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
	var input = tools.ConvertToInt(tools.Read(args[0]))

	var prev = input[0]
	var count = 0
	for _, v := range input[1:] {
		if prev < v {
			count++
		}
		prev = v
	}

	fmt.Println("Increasing Count: ", count)

	return 0
}

func part2(args []string) int {
	var input = tools.ConvertToInt(tools.Read(args[0]))

	var prev = input[0] + input[1] + input[2]
	var count = 0
	fmt.Println(prev)
	for i := 1; i < len(input)-2; i++ {
		fmt.Print(input[i], input[i+1], input[i+2])
		var window = input[i] + input[i+1] + input[i+2]
		if window > prev {
			fmt.Printf(" = %d increased\n", window)
			count++
		} else {
			fmt.Printf(" = %d\n", window)
		}

		prev = window
	}

	fmt.Println("Increasing Count: ", count)

	return 0
}
