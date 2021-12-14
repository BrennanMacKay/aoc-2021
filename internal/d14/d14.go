package d14

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"math"
	"regexp"
	"strconv"
)

type pair struct {
	first  string
	second string
}

type result struct {
	first  pair
	second pair
}

var rRegex = regexp.MustCompile(`(.)(.) -> (.)`)

func Day14(args []string) int {
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
	pairs, last := readTemplate(input[0])
	rules := readRules(input[2:])

	fmt.Println(pairs)
	fmt.Println(rules)

	steps, _ := strconv.Atoi(args[1])
	for i := 0; i < steps; i++ {
		pairs = step(rules, pairs)
	}

	fmt.Println(pairs, last)

	counts := count(pairs, last)

	fmt.Println(counts)

	min, max := minMax(counts)

	fmt.Println(min, max)
	fmt.Println(max - min)

	return 0
}

func part2(args []string) int {
	return 0
}

func minMax(counts map[string]int) (int, int) {
	min, max := math.MaxInt, 0
	for _, c := range counts {
		if c > max {
			max = c
		}
		if c < min {
			min = c
		}
	}

	return min, max
}

func count(pairs map[pair]int, last string) map[string]int {
	counts := make(map[string]int)

	for p, c := range pairs {
		counts[p.first] += c
	}

	counts[last]++

	return counts
}

func step(rules map[pair]result, pairs map[pair]int) map[pair]int {
	newPairs := make(map[pair]int)
	for pair, count := range pairs {
		result := rules[pair]
		newPairs[result.first] += count
		newPairs[result.second] += count
	}

	return newPairs
}

func readTemplate(input string) (map[pair]int, string) {
	template := []rune(input)
	pairCounts := make(map[pair]int)
	for i := 0; i < len(template)-1; i++ {
		pairCounts[pair{string(template[i]), string(template[i+1])}] += 1
	}

	return pairCounts, string(template[len(template)-1])
}

// expects rules to start on first line
func readRules(input []string) map[pair]result {
	rules := make(map[pair]result, len(input))
	for _, rule := range input {
		parts := rRegex.FindStringSubmatch(rule)
		p1, p2, r := parts[1], parts[2], parts[3]
		p := pair{p1, p2}
		res := result{pair{p.first, r}, pair{r, p.second}}

		rules[p] = res
	}

	return rules
}
