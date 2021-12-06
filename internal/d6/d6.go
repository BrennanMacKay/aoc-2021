package d6

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"strconv"
	"strings"
)

const daysInCycle = 8
const initialDelay = 2

func Day6(args []string) int {
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
	fishDays := strings.Split(tools.Read(args[0])[0], ",")
	fish := buildFishMap(fishDays)

	fmt.Println(fish)
	daysToRun, _ := strconv.Atoi(args[1])

	for i := 1; i <= daysToRun; i++ {
		fish = processDay(fish)
		fmt.Println("Day", i, fish)
	}

	fmt.Println(sum(fish))

	return 0
}

func part2(args []string) int {
	return 0
}

func sum(fishDays map[int]int) int {
	var acc int
	for _, c := range fishDays {
		acc += c
	}
	return acc
}

func processDay(fishDays map[int]int) map[int]int {
	nextDay := make(map[int]int, daysInCycle)
	for day, count := range fishDays {
		if day == 0 {
			nextDay[daysInCycle] += count
			nextDay[daysInCycle-initialDelay] += count
		} else {
			nextDay[day-1] += count
		}
	}

	return nextDay
}

// we could use an array instead, but a map is easier to debug
func buildFishMap(fishDays []string) map[int]int {
	fishMap := make(map[int]int, daysInCycle)
	for _, fish := range fishDays {
		days, _ := strconv.Atoi(fish)
		fishMap[days]++
	}

	return fishMap
}
