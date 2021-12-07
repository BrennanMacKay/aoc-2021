package d7

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"math"
	"sort"
	"strconv"
	"strings"
)

func Day7(args []string) int {
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
	input := strings.Split(tools.Read(args[0])[0], ",")
	positions := make([]int, len(input))
	for i, s := range input {
		positions[i], _ = strconv.Atoi(s)
	}

	dest := median(positions)
	fmt.Println("MED", dest)
	fmt.Println("Fuel", calcFuelP1(positions, dest))

	return 0
}

func part2(args []string) int {
	input := strings.Split(tools.Read(args[0])[0], ",")
	positions := make([]int, len(input))
	for i, s := range input {
		positions[i], _ = strconv.Atoi(s)
	}

	dest := avg(positions)
	fmt.Println("MED", dest)
	fmt.Println("Fuel", calcFuelP2(positions, dest))
	fmt.Println("Fuel + 1", calcFuelP2(positions, dest+1))
	fmt.Println("Fuel - 1", calcFuelP2(positions, dest-1))

	return 0
}

func calcFuelP1(positions []int, dest int) int {
	totalFuel := 0
	fmt.Println("Destination", dest)
	for _, position := range positions {
		fuel := int(math.Abs(float64(dest - position)))
		totalFuel += fuel
		//fmt.Println("Position", position, "Fuel Req", fuel)
	}

	return totalFuel
}

func calcFuelP2(positions []int, dest int) int {
	totalFuel := 0
	fmt.Println("Destination", dest)
	for _, position := range positions {
		distance := int(math.Abs(float64(dest - position)))
		fuel := distanceToFuel(distance)
		totalFuel += fuel
		//fmt.Println("Position", position, "Fuel Req", fuel)
	}

	return totalFuel
}

func distanceToFuel(dist int) int {
	acc := 0
	for i := dist; i > 0; i-- {
		acc += i
	}

	return acc
}

func median(values []int) int {
	sort.Ints(values)
	mid := len(values) / 2

	if len(values)%2 == 0 {
		fmt.Println(values[mid-1], values[mid])
		return (values[mid-1] + values[mid]) / 2
	} else {
		return values[mid]
	}
}

func avg(values []int) int {
	acc := 0
	for _, v := range values {
		acc += v
	}

	avg := float64(acc) / float64(len(values))
	return int(math.Round(avg))
}
