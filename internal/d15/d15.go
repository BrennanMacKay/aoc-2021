package d15

import (
	"container/heap"
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"math"
	"strconv"
)

// This is all terrible!

type Item struct {
	Node     Node
	Priority int
	Index    int
}

type Point struct {
	x int
	y int
}

type Node struct {
	tDist int // should be set to max int
	lDist int // 1 - 9
	point Point
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = i
	pq[j].Index = i
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.Index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *Item, value Node, priority int) {
	item.Node = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}

func Day15(args []string) int {
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
	grid := readInput(input)
	pq, items := addToQueue(grid)
	unvisited := unvisited(grid)
	//printQueue(pq)
	result := dijkstras(&grid, pq, unvisited, items)

	print(grid)
	fmt.Println(result)

	return 0
}

func part2(args []string) int {
	input := tools.Read(args[0])
	grid := readInput(input)
	grid = expandGrid(grid)
	//print(grid)
	pq, items := addToQueue(grid)
	unvisited := unvisited(grid)
	//printQueue(pq)
	result := dijkstras(&grid, pq, unvisited, items)

	//print(grid)
	fmt.Println(result)

	return 0
}

func printQueue(pq *PriorityQueue) {
	for i := 0; i < pq.Len(); i++ {
		fmt.Println((*pq)[i])
	}
}

func dijkstras(grid *[][]Node, pq *PriorityQueue, unvisited map[Point]struct{}, items map[Point]*Item) int {
	current := (*grid)[0][0]
	max := Point{len((*grid)[0]), len(*grid)}

	for pq.Len() > 0 {
		if pq.Len()%1000 == 0 {
			fmt.Println(pq.Len())
		}
		neighbors := neighbors(current.point, max, unvisited)
		//fmt.Println(current.point, "from", neighbors)
		if current.tDist == math.MaxInt {
			fmt.Println("WOOPS!")
			break
		}

		for n := range neighbors {
			node := (*grid)[n.y][n.x]
			tDist := node.lDist + current.tDist
			if tDist < node.tDist {
				node.tDist = tDist
				(*grid)[n.y][n.x] = node
				item := items[node.point]
				pq.Update(item, node, tDist)
				//fmt.Println("UPDATE", node.point, tDist)
			}
		}
		delete(unvisited, current.point)
		heap.Init(pq)
		current = heap.Pop(pq).(*Item).Node
		//printQueue(pq)
	}

	dest := (*grid)[len(*grid)-1][len((*grid)[0])-1]
	return dest.tDist

}

func neighbors(current Point, max Point, unvisited map[Point]struct{}) map[Point]struct{} {
	points := make(map[Point]struct{})
	if current.x > 0 {
		points[Point{current.x - 1, current.y}] = struct{}{}
	}
	if current.x < max.x {
		points[Point{current.x + 1, current.y}] = struct{}{}
	}
	if current.y > 0 {
		points[Point{current.x, current.y - 1}] = struct{}{}
	}
	if current.y < max.y {
		points[Point{current.x, current.y + 1}] = struct{}{}
	}

	pointsToRemove := make([]Point, 0)

	for p := range points {
		if _, ok := unvisited[p]; !ok {
			pointsToRemove = append(pointsToRemove, p)
		}
	}

	for _, p := range pointsToRemove {
		delete(points, p)
	}

	return points
}

func print(grid [][]Node) {
	for _, row := range grid {
		for _, node := range row {
			fmt.Printf("%d", node.lDist)
		}
		fmt.Println()
	}
}

func unvisited(grid [][]Node) map[Point]struct{} {
	unvisited := make(map[Point]struct{}, (len(grid)*len(grid[0]))-1)

	for y, row := range grid {
		for x, node := range row {
			if x == 0 && y == 0 {
				// skip the first item since this is where we start
				continue
			}

			unvisited[node.point] = struct{}{}
		}
	}

	return unvisited
}

func addToQueue(grid [][]Node) (*PriorityQueue, map[Point]*Item) {
	pq := make(PriorityQueue, (len(grid)*len(grid[0]))-1)
	items := make(map[Point]*Item)

	i := 0
	for y, row := range grid {
		for x, node := range row {
			if x == 0 && y == 0 {
				// skip the first item since this is where we start
				continue
			}
			item := Item{Node: node, Priority: node.tDist, Index: i}
			pq[i] = &item
			items[node.point] = &item
			i++
		}
	}

	return &pq, items
}

func expandGrid(grid [][]Node) [][]Node {
	ySize := len(grid)
	xSize := len(grid[0])

	newGrid := make([][]Node, ySize*5)
	for y := range newGrid {
		newGrid[y] = make([]Node, xSize*5)
	}

	for y, rows := range newGrid {
		ySource := y % ySize
		yScale := y / ySize
		for x := range rows {
			xSource := x % xSize
			xScale := x / xSize
			scale := xScale + yScale
			source := grid[ySource][xSource]
			newRisk := source.lDist + scale
			if newRisk >= 10 {
				newRisk -= 9
			}
			tDisk := math.MaxInt
			if x == 0 && y == 0 {
				tDisk = 0
			}
			newGrid[y][x] = Node{tDist: tDisk, lDist: newRisk, point: Point{x, y}}
		}
	}

	return newGrid
}

func readInput(input []string) [][]Node {
	grid := make([][]Node, len(input))

	for y := range grid {
		grid[y] = make([]Node, len(input[0]))
	}

	for y, row := range input {
		for x, r := range []rune(row) {
			tDist := math.MaxInt
			risk, _ := strconv.Atoi(string(r))
			if x == 0 && y == 0 {
				tDist = 0
			}

			point := Point{x, y}
			node := Node{tDist, risk, point}
			grid[y][x] = node
		}
	}

	return grid
}
