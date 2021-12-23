package d22

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"regexp"
)

type Instruction struct {
	on   bool
	cube Cuboid
}

type Point struct {
	x int
	y int
	z int
}

type Cuboid struct {
	lBound Point
	uBound Point
}

type state map[Cuboid]struct{}
type simpleState map[Point]struct{}

var reg = regexp.MustCompile(`(\w*) x=(-?\d*)\.\.(-?\d*),y=(-?\d*)\.\.(-?\d*),z=(-?\d*)\.\.(-?\d*)`)

func Day22(args []string) int {
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
	lowerBound := Point{-50, -50, -50}
	upperBound := Point{51, 51, 51}
	bound := Cuboid{lowerBound, upperBound}

	instructs := readInput(tools.Read(args[0]))
	instructs = restrictToBounds(instructs, bound)

	fmt.Println(instructs)

	simpleState := make(simpleState)
	for i, ins := range instructs {
		for z := ins.cube.lBound.z; z < ins.cube.uBound.z; z++ {
			for y := ins.cube.lBound.y; y < ins.cube.uBound.y; y++ {
				for x := ins.cube.lBound.x; x < ins.cube.uBound.x; x++ {
					if ins.on {
						simpleState[Point{x, y, z}] = struct{}{}
					} else {
						delete(simpleState, Point{x, y, z})
					}
				}
			}
		}
		fmt.Println(i, len(simpleState))
	}

	fmt.Println(len(simpleState))

	return 0
}

func part2(args []string) int {
	instructs := readInput(tools.Read(args[0]))
	//lowerBound := Point{-60, -60, -60}
	//upperBound := Point{60, 60, 60}
	//bound := Cuboid{lowerBound, upperBound}
	//
	//instructs = restrictToBounds(instructs, bound)

	for i, v := range instructs {
		fmt.Println(i, v)
	}

	state := make(state)
	for i, ins := range instructs {
		fmt.Println("****", i)
		//fmt.Println("BEFORE:", state, count(state))
		process(ins, state)
		fmt.Println(i, ins, len(state))
		//fmt.Println("AFTER:", state)
		fmt.Println("COUNT:", count(state))
		if i == 1 {
			//break
		}
	}

	res := count(state)

	fmt.Println(res)

	return 0
}

func process(ins Instruction, state state) {
	if ins.on {
		ins.cube.on(state)
	} else {
		ins.cube.off(state)
	}
}

func count(state state) int {
	acc := 0
	for c := range state {
		acc += c.size()
	}

	return acc
}

func readInput(input []string) []Instruction {
	instructs := make([]Instruction, len(input))

	for i, v := range input {
		parts := reg.FindStringSubmatch(v)
		lower := Point{tools.ToInt(parts[2]), tools.ToInt(parts[4]), tools.ToInt(parts[6])}
		upper := Point{tools.ToInt(parts[3]) + 1, tools.ToInt(parts[5]) + 1, tools.ToInt(parts[7]) + 1}
		cube := Cuboid{lower, upper}
		instructs[i] = Instruction{parts[1] == "on", cube}
	}

	return instructs
}

func restrictToBounds(instructions []Instruction, bound Cuboid) []Instruction {
	newInstructs := make([]Instruction, 0)
	for _, v := range instructions {
		c := v.cube
		if c.within(bound) {
			newInstructs = append(newInstructs, v)
		}
	}

	return newInstructs
}

func (c Cuboid) on(s state) {
	additions := make([]Cuboid, 0)
	removals := make([]Cuboid, 0)

	for other := range s {
		if inter, yes := c.Intersect(other); yes {
			fmt.Println("INTERSECTION ON:", "THIS:", c, "OTHER:", other, "INTER:", inter)
			removals = append(removals, other)
			cubes := other.split(inter)
			for _, cube := range cubes {
				additions = append(additions, cube)
			}
		}
	}

	for _, add := range additions {

		if _, ok := s[add]; ok {
			fmt.Println("\t\t**** EXISTS ***", add)
		}
		s[add] = struct{}{}
	}

	for _, rem := range removals {
		delete(s, rem)
	}

	s[c] = struct{}{}
}

// similar to on but we don't add this to the state
func (c Cuboid) off(s state) {
	additions := make([]Cuboid, 0)
	removals := make([]Cuboid, 0)

	for other := range s {
		if inter, yes := c.Intersect(other); yes {
			fmt.Println("INTERSECTION OFF", "THIS:", c, "OTHER:", other, "INTER:", inter)
			removals = append(removals, other)
			cubes := other.split(inter)
			for _, cube := range cubes {
				additions = append(additions, cube)
			}
		}
	}

	for _, add := range additions {
		s[add] = struct{}{}
	}

	for _, rem := range removals {
		delete(s, rem)
	}

}

func (c Cuboid) size() int {
	x, y, z := c.uBound.x-c.lBound.x, c.uBound.y-c.lBound.y, c.uBound.z-c.lBound.z

	return x * y * z
}

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func (c Cuboid) split(inter Cuboid) []Cuboid {
	newCubes := make([]Cuboid, 0)

	lxCube := Cuboid{Point{c.lBound.x, c.lBound.y, c.lBound.z}, Point{inter.lBound.x, c.uBound.y, c.uBound.z}}
	uxCube := Cuboid{Point{inter.uBound.x, c.lBound.y, c.lBound.z}, Point{c.uBound.x, c.uBound.y, c.uBound.z}}

	lyCube := Cuboid{Point{lxCube.uBound.x, c.lBound.y, c.lBound.z}, Point{uxCube.lBound.x, inter.lBound.y, c.uBound.z}}
	uyCube := Cuboid{Point{lxCube.uBound.x, inter.uBound.y, c.lBound.z}, Point{uxCube.lBound.x, c.uBound.y, c.uBound.z}}

	lzCube := Cuboid{Point{lxCube.uBound.x, lyCube.uBound.y, c.lBound.z}, Point{uxCube.lBound.x, uyCube.lBound.y, inter.lBound.z}}
	uzCube := Cuboid{Point{lxCube.uBound.x, lyCube.uBound.y, inter.uBound.z}, Point{uxCube.lBound.x, uyCube.lBound.y, c.uBound.z}}

	fmt.Println("\tX", lxCube, uxCube)
	fmt.Println("\tY", lyCube, uyCube)
	fmt.Println("\tZ", lzCube, uzCube)

	if lxCube.size() > 0 {
		newCubes = append(newCubes, lxCube)
	}
	if uxCube.size() > 0 {
		newCubes = append(newCubes, uxCube)
	}
	if lyCube.size() > 0 {
		newCubes = append(newCubes, lyCube)
	}
	if uyCube.size() > 0 {
		newCubes = append(newCubes, uyCube)
	}
	if lzCube.size() > 0 {
		newCubes = append(newCubes, lzCube)
	}
	if uzCube.size() > 0 {
		newCubes = append(newCubes, uzCube)
	}

	fmt.Println("\tTHIS", c, "OTHER", inter)
	fmt.Println("\tNEW:", newCubes)

	return newCubes
}

func (p Point) withinLower(other Point) bool {
	return p.x >= other.x && p.y >= other.y && p.z >= other.z
}

func (p Point) withinLowerGT(other Point) bool {
	return p.x > other.x && p.y > other.y && p.z > other.z
}

func (p Point) withinUpper(other Point) bool {
	return p.x < other.x && p.y < other.y && p.z < other.z
}

func (p Point) within(otherLower Point, otherUpper Point) bool {
	return p.withinLower(otherLower) && p.withinUpper(otherUpper)
}

func (c Cuboid) within(other Cuboid) bool {
	return c.lBound.withinLower(other.lBound) && c.uBound.withinUpper(other.uBound)
}

func (c Cuboid) Intersect(other Cuboid) (Cuboid, bool) {
	if c.within(other) {
		fmt.Println("\t\tc within")
		return c, true
	} else if other.within(c) {
		fmt.Println("\t\tother within")
		return other, true
	} else if !c.lBound.withinUpper(other.uBound) || !c.uBound.withinLowerGT(other.lBound) {
		return Cuboid{}, false
	}

	var newLower Point
	var newUpper Point
	if c.lBound.withinLower(other.lBound) && c.lBound.withinUpper(other.uBound) {
		newLower = c.lBound
	} else if other.lBound.withinLower(c.lBound) && other.lBound.withinUpper(c.uBound) {
		newLower = other.lBound
	} else {
		// this is awful :(
		if c.lBound.x >= other.lBound.x && c.lBound.x < other.uBound.x {
			newLower.x = c.lBound.x
		} else {
			newLower.x = other.lBound.x
		}

		if c.lBound.y >= other.lBound.y && c.lBound.y < other.uBound.y {
			newLower.y = c.lBound.y
		} else {
			newLower.y = other.lBound.y
		}

		if c.lBound.z >= other.lBound.z && c.lBound.z < other.uBound.z {
			newLower.z = c.lBound.z
		} else {
			newLower.z = other.lBound.z
		}
	}

	if c.uBound.withinLowerGT(other.lBound) && c.uBound.withinUpper(other.uBound) {
		newUpper = c.uBound
	} else if other.uBound.withinLowerGT(c.lBound) && other.uBound.withinUpper(c.uBound) {
		newUpper = other.uBound
	} else {
		if c.uBound.x > other.lBound.x && c.uBound.x <= other.uBound.x {
			newUpper.x = c.uBound.x
		} else {
			newUpper.x = other.uBound.x
		}

		if c.uBound.y > other.lBound.y && c.uBound.y <= other.uBound.y {
			newUpper.y = c.uBound.y
		} else {
			newUpper.y = other.uBound.y
		}

		if c.uBound.z > other.lBound.z && c.uBound.z <= other.uBound.z {
			newUpper.z = c.uBound.z
		} else {
			newUpper.z = other.uBound.z
		}
	}

	return Cuboid{newLower, newUpper}, true
}

func (c Cuboid) printPoints() {
	for z := c.lBound.z; z < c.uBound.z; z++ {
		for y := c.lBound.y; y < c.uBound.y; y++ {
			for x := c.lBound.x; x < c.uBound.x; x++ {
				fmt.Printf("%d, %d, %d", x, y, z)
			}
		}
	}
}

//func (s state) printState() {
//	points := make([]Point, 0)
//	for k := range s {
//		points = append(points, k)
//	}
//
//	sort.Slice(points, func(i, j int) bool {
//		x := points[i].x < points[j].x
//		y := points[i].y < points[j].y
//		z := points[i].z < points[j].z
//		return z || y || x
//	})
//
//	for _, v := range points {
//		fmt.Println(v)
//	}
//}
