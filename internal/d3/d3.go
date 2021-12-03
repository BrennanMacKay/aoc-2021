package d3

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"strconv"
)

// This is really awful! Look away!

func Day3(args []string) int {
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
	var input = tools.Read(args[0])

	columns := makeColumns(input)
	fmt.Printf("%v\n", columns)

	counts := makeCounts(columns)

	var gammaRate string
	var epsilonRate string
	for _, count := range counts {
		if count > len(input)/2 {
			gammaRate += "1"
			epsilonRate += "0"
		} else {
			gammaRate += "0"
			epsilonRate += "1"
		}
	}

	fmt.Println("Gamma", gammaRate, "Epsilon", epsilonRate)

	var gamma, _ = strconv.ParseInt(gammaRate, 2, 0)
	var epsilon, _ = strconv.ParseInt(epsilonRate, 2, 0)
	fmt.Println(gamma * epsilon)

	return 0
}

func part2(args []string) int {
	var input = tools.Read(args[0])
	fmt.Println(input)

	columns := makeColumns(input)
	fmt.Printf("%v\n", columns)

	var oxyIndex = oxygenIndex(columns)
	var co2Index = co2Index(columns)

	fmt.Println("oxy", input[oxyIndex])
	fmt.Println("co2", input[co2Index])

	var oxy, _ = strconv.ParseInt(input[oxyIndex], 2, 0)
	var co2, _ = strconv.ParseInt(input[co2Index], 2, 0)

	fmt.Println(oxy * co2)

	return 0
}

func oxygenIndex(columns [][]int) int {
	var validCodes = make(map[int]struct{})
	for i := range columns[0] {
		validCodes[i] = struct{}{}
	}

	for i, column := range columns {
		fmt.Printf("%v\n", validCodes)
		var ones int
		for j, bit := range column {
			if _, ok := validCodes[j]; ok {
				ones += bit
			}
		}
		zeros := len(validCodes) - ones

		var req int
		fmt.Println("Count", ones, "col", i, "Codes", len(validCodes))
		if ones >= zeros {
			req = 1
		} else {
			req = 0
		}

		for j, bit := range column {
			if _, ok := validCodes[j]; ok {
				if bit != req {
					fmt.Println(bit, "Not Valid as looking for", req, "in item", j, "col", i)
					delete(validCodes, j)
				}
			}
		}
		fmt.Printf("%v\n", validCodes)

		if len(validCodes) < 2 {
			for k := range validCodes {
				return k
			}
		}
	}

	fmt.Printf("last %v\n", validCodes)

	return -1
}

func co2Index(columns [][]int) int {
	var validCodes = make(map[int]struct{})
	for i := range columns[0] {
		validCodes[i] = struct{}{}
	}

	for i, column := range columns {
		fmt.Printf("%v\n", validCodes)
		var ones int
		for j, bit := range column {
			if _, ok := validCodes[j]; ok {
				ones += bit
			}
		}
		zeros := len(validCodes) - ones

		var req int
		fmt.Println("Count", ones, "col", i, "Codes", len(validCodes))
		if ones < zeros {
			req = 1
		} else {
			req = 0
		}

		for j, bit := range column {
			if _, ok := validCodes[j]; ok {
				if bit != req {
					fmt.Println(bit, "Not Valid as looking for", req, "in item", j, "col", i)
					delete(validCodes, j)
				}
			}
		}
		fmt.Printf("%v\n", validCodes)

		if len(validCodes) < 2 {
			for k := range validCodes {
				return k
			}
		}
	}

	fmt.Printf("last %v\n", validCodes)

	return -1
}

func makeColumns(input []string) [][]int {
	columns := make([][]int, len(input[0]))

	for i := range columns {
		columns[i] = make([]int, len(input))
	}

	for i, code := range input {
		for j, bit := range []rune(code) {
			columns[j][i] = int(bit) - 48
		}
	}

	return columns
}

func makeCounts(columns [][]int) []int {
	counts := make([]int, len(columns))
	for i, column := range columns {
		for _, bit := range column {
			counts[i] += bit
		}
	}

	return counts
}
