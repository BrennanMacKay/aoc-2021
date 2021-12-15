package main

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/d0"
	"github.com/BrennanMacKay/aoc-2021/internal/d1"
	"github.com/BrennanMacKay/aoc-2021/internal/d10"
	"github.com/BrennanMacKay/aoc-2021/internal/d11"
	"github.com/BrennanMacKay/aoc-2021/internal/d12"
	"github.com/BrennanMacKay/aoc-2021/internal/d13"
	"github.com/BrennanMacKay/aoc-2021/internal/d14"
	"github.com/BrennanMacKay/aoc-2021/internal/d15"
	"github.com/BrennanMacKay/aoc-2021/internal/d2"
	"github.com/BrennanMacKay/aoc-2021/internal/d3"
	"github.com/BrennanMacKay/aoc-2021/internal/d4"
	"github.com/BrennanMacKay/aoc-2021/internal/d5"
	"github.com/BrennanMacKay/aoc-2021/internal/d6"
	"github.com/BrennanMacKay/aoc-2021/internal/d7"
	"github.com/BrennanMacKay/aoc-2021/internal/d8"
	"github.com/BrennanMacKay/aoc-2021/internal/d9"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Provide d{day} p{part} sub commands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "d0":
		os.Exit(d0.Day0(os.Args[2:]))
	case "d1":
		os.Exit(d1.Day1(os.Args[2:]))
	case "d2":
		os.Exit(d2.Day2(os.Args[2:]))
	case "d3":
		os.Exit(d3.Day3(os.Args[2:]))
	case "d4":
		os.Exit(d4.Day4(os.Args[2:]))
	case "d5":
		os.Exit(d5.Day5(os.Args[2:]))
	case "d6":
		os.Exit(d6.Day6(os.Args[2:]))
	case "d7":
		os.Exit(d7.Day7(os.Args[2:]))
	case "d8":
		os.Exit(d8.Day8(os.Args[2:]))
	case "d9":
		os.Exit(d9.Day9(os.Args[2:]))
	case "d10":
		os.Exit(d10.Day10(os.Args[2:]))
	case "d11":
		os.Exit(d11.Day11(os.Args[2:]))
	case "d12":
		os.Exit(d12.Day12(os.Args[2:]))
	case "d13":
		os.Exit(d13.Day13(os.Args[2:]))
	case "d14":
		os.Exit(d14.Day14(os.Args[2:]))
	case "d15":
		os.Exit(d15.Day15(os.Args[2:]))
	default:
		fmt.Printf("%s did not match a known problem\n", os.Args[1])
		os.Exit(1)
	}
}
