package d12

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"strings"
)

type cave struct {
	name    string
	isLarge bool
	adj     map[string]struct{}
}

var caves map[string]cave

type route []string

func Day12(args []string) int {
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
	readCaves(input)
	res := routesFrom("start", make(route, 0), make(map[string]int), false)

	printCaves()
	fmt.Println(res)
	fmt.Println(len(res))

	return 0
}

func part2(args []string) int {
	input := tools.Read(args[0])
	readCaves(input)
	res := routesFrom("start", make(route, 0), make(map[string]int), true)

	printCaves()
	//fmt.Println(res)
	fmt.Println(len(res))

	return 0
}

func routesFrom(c string, r route, visited map[string]int, canDouble bool) []route {
	//time.Sleep(time.Millisecond)
	//fmt.Println(r, c, canDouble, visited)
	r = append(r, c)
	if c == "end" {
		//fmt.Println("ROUTE", r)
		return []route{r}
	}

	curr := caves[c]
	visited[curr.name]++

	routes := make([]route, 0)
	for adj := range curr.adj {
		if adj == "start" {
			continue
		} else if count, ok := visited[adj]; ok && count > 0 && strings.ToLower(adj) == adj {
			if canDouble {
				res := routesFrom(adj, r, visited, false)
				routes = append(routes, res...)
				visited[adj]--
			} else {
				continue
			}
		} else {
			res := routesFrom(adj, r, visited, canDouble)
			routes = append(routes, res...)
			visited[adj]--
		}
	}

	return routes
}

func readCaves(input []string) {
	caves = make(map[string]cave, len(input))
	for _, path := range input {
		parts := strings.Split(path, "-")
		addPath(parts[0], parts[1])
	}
}

func addPath(this string, other string) {

	thisCave, ok := caves[this]
	if !ok {
		thisCave = cave{this, strings.ToUpper(this) == this, make(map[string]struct{}, 0)}
		caves[this] = thisCave
	}

	otherCave, ok := caves[other]
	if !ok {
		otherCave = cave{other, strings.ToUpper(other) == other, make(map[string]struct{}, 0)}
		caves[other] = otherCave
	}

	thisCave.adj[otherCave.name] = struct{}{}
	otherCave.adj[thisCave.name] = struct{}{}

	caves[this] = thisCave
	caves[other] = otherCave

}

func printCaves() {
	for k, v := range caves {
		fmt.Println(k, v)
	}
}
