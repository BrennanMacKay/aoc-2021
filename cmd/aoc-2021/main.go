package main

import (
	"fmt"
	"os"

	"github.com/BrennanMacKay/aoc-2021/internal/problems"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Provide d{day} p{part} sub commands")
		os.Exit(1)
	}


	switch os.Args[1] {
	case "d0":
		os.Exit(problems.Day0(os.Args[2:]))
	default:
		fmt.Printf("%s did not match a known problem\n", os.Args[1])
		os.Exit(1)
	}
}
