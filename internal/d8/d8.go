package d8

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"regexp"
	"strconv"
)

var regx = regexp.MustCompile(`(\w*) (\w*) (\w*) (\w*) (\w*) (\w*) (\w*) (\w*) (\w*) (\w*) \| (\w*) (\w*) (\w*) (\w*)`)

type instance struct {
	key     [10]string
	display [4]string

	displayMap [4]map[string]struct{}

	one   map[string]struct{}
	four  map[string]struct{}
	seven map[string]struct{}
	eight map[string]struct{}
}

func Day8(args []string) int {
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
	instances := parseInput(input)

	acc := 0
	for _, inst := range instances {
		for _, v := range inst.display {
			l := len(v)
			if l == 2 || l == 3 || l == 4 || l == 7 {
				acc++
			}
		}
	}

	fmt.Println(acc)

	return 0
}

func part2(args []string) int {
	input := tools.Read(args[0])
	instances := parseInput(input)

	results := make([]string, len(instances))
	for rIndex, inst := range instances {
		inst = buildDisplayValues(inst)
		inst = populateKnown(inst)
		result := ""
		for i, display := range inst.display {
			switch len(display) {
			case 2:
				result += "1"
			case 3:
				result += "7"
			case 4:
				result += "4"
			case 5: // 2, 3, 5
				if contains(inst.displayMap[i], inst.seven) {
					result += "3"
				} else {
					intersection := intersects(inst.displayMap[i], inst.four)
					if len(intersection) == 3 {
						result += "5"
					} else {
						result += "2"
					}
				}
			case 6: // 0, 6, 9
				if contains(inst.displayMap[i], inst.four) {
					result += "9"
				} else if contains(inst.displayMap[i], inst.seven) {
					result += "0"
				} else {
					result += "6"
				}
			case 7:
				result += "8"
			}
		}

		results[rIndex] = result
	}

	fmt.Println(results)

	sum := 0
	for _, r := range results {
		v, _ := strconv.Atoi(r)
		sum += v
	}

	fmt.Println(sum)
	return 0
}

func buildDisplayValues(inst instance) instance {
	for i, v := range inst.display {
		inst.displayMap[i] = make(map[string]struct{}, len(v))
		for _, r := range []rune(v) {
			inst.displayMap[i][string(r)] = struct{}{}
		}
	}

	return inst
}

func populateKnown(inst instance) instance {
	for _, v := range inst.key {
		switch len(v) {
		case 2: //1
			inst.one = make(map[string]struct{}, len(v))
			for _, r := range []rune(v) {
				inst.one[string(r)] = struct{}{}
			}
		case 3: //7
			inst.seven = make(map[string]struct{}, len(v))
			for _, r := range []rune(v) {
				inst.seven[string(r)] = struct{}{}
			}
		case 4: //4
			inst.four = make(map[string]struct{}, len(v))
			for _, r := range []rune(v) {
				inst.four[string(r)] = struct{}{}
			}
		case 7: //8
			inst.eight = make(map[string]struct{}, len(v))
			for _, r := range []rune(v) {
				inst.eight[string(r)] = struct{}{}
			}
		}
	}

	return inst
}

func contains(this map[string]struct{}, that map[string]struct{}) bool {
	for k := range that {
		if _, ok := this[k]; !ok {
			return false
		}
	}

	return true
}

func intersects(this map[string]struct{}, that map[string]struct{}) map[string]struct{} {
	intersection := make(map[string]struct{}, len(this))
	for k := range that {
		if v, ok := this[k]; ok {
			intersection[k] = v
		}
	}
	return intersection
}

func parseInput(input []string) []instance {
	instances := make([]instance, len(input))
	for i, v := range input {
		instances[i] = parseLine(v)
	}

	return instances
}

func parseLine(input string) instance {
	sParts := regx.FindStringSubmatch(input)
	inst := instance{}
	for i, s := range sParts[1:11] {
		inst.key[i] = s
	}
	for i, s := range sParts[11:15] {
		inst.display[i] = s
	}

	return inst
}
