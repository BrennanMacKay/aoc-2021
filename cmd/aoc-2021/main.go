package main

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/d0"
	"github.com/BrennanMacKay/aoc-2021/internal/d1"
	"github.com/BrennanMacKay/aoc-2021/internal/d2"
	"github.com/BrennanMacKay/aoc-2021/internal/d3"
	"github.com/BrennanMacKay/aoc-2021/internal/d4"
	"github.com/BrennanMacKay/aoc-2021/internal/d5"
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
	default:
		fmt.Printf("%s did not match a known problem\n", os.Args[1])
		os.Exit(1)
	}
}
