package d22

import (
	"fmt"
	"testing"
)

func TestIntersectCubeWithStick(t *testing.T) {
	square := Cuboid{p(9, 9, 9), p(12, 12, 12)}

	var tests = []struct {
		a, b, want Cuboid
		inter      bool
	}{
		{square, Cuboid{p(8, 10, 10), p(11, 11, 11)}, Cuboid{p(9, 10, 10), p(11, 11, 11)}, true},
		{square, Cuboid{p(10, 10, 10), p(13, 11, 11)}, Cuboid{p(10, 10, 10), p(12, 11, 11)}, true},
		{square, Cuboid{p(10, 8, 10), p(11, 11, 11)}, Cuboid{p(10, 9, 10), p(11, 11, 11)}, true},
		{square, Cuboid{p(10, 10, 10), p(11, 13, 11)}, Cuboid{p(10, 10, 10), p(11, 12, 11)}, true},
		{square, Cuboid{p(10, 10, 8), p(11, 11, 11)}, Cuboid{p(10, 10, 9), p(11, 11, 11)}, true},
		{square, Cuboid{p(10, 10, 10), p(11, 11, 13)}, Cuboid{p(10, 10, 10), p(11, 11, 12)}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			c, yes := tt.a.Intersect(tt.b)
			if !yes || c != tt.want {
				t.Errorf("got %v %v, want %v", yes, c, tt.want)
			}
		})
		testname = fmt.Sprintf("%v, %v", tt.b, tt.a)
		t.Run(testname, func(t *testing.T) {
			c, yes := tt.b.Intersect(tt.a)
			if !yes || c != tt.want {
				t.Errorf("got %v %v, want %v", yes, c, tt.want)
			}
		})

	}
}

func TestIntersectCubeWithCorner(t *testing.T) {
	square := Cuboid{p(9, 9, 9), p(11, 11, 11)}

	var tests = []struct {
		a, b, want Cuboid
		inter      bool
	}{
		{square, Cuboid{p(8, 8, 8), p(10, 10, 10)}, Cuboid{p(9, 9, 9), p(10, 10, 10)}, true},
		{square, Cuboid{p(10, 8, 8), p(12, 10, 10)}, Cuboid{p(10, 9, 9), p(11, 10, 10)}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			c, yes := tt.a.Intersect(tt.b)
			if !yes || c != tt.want {
				t.Errorf("got %v %v, want %v", yes, c, tt.want)
			}
		})
		testname = fmt.Sprintf("%v, %v", tt.b, tt.a)
		t.Run(testname, func(t *testing.T) {
			c, yes := tt.b.Intersect(tt.a)
			if !yes || c != tt.want {
				t.Errorf("got %v %v, want %v", yes, c, tt.want)
			}
		})

	}
}

func TestIntersectStrange(t *testing.T) {
	square := Cuboid{p(-18, -33, -7), p(27, 16, 47)}

	var tests = []struct {
		a, b, want Cuboid
		inter      bool
	}{
		{square, Cuboid{p(6, -27, 20), p(22, -8, 30)}, Cuboid{p(6, -27, 20), p(22, -8, 30)}, true},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%v, %v", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			c, yes := tt.a.Intersect(tt.b)
			if !yes || c != tt.want {
				t.Errorf("got %v %v, want %v", yes, c, tt.want)
			}
		})
		testname = fmt.Sprintf("%v, %v", tt.b, tt.a)
		t.Run(testname, func(t *testing.T) {
			c, yes := tt.b.Intersect(tt.a)
			if !yes || c != tt.want {
				t.Errorf("got %v %v, want %v", yes, c, tt.want)
			}
		})

	}
}

func p(x, y, z int) Point {
	return Point{x, y, z}
}
