package d10

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"sort"
)

type stack []rune

var closeMatch = map[rune]rune{')': '(', ']': '[', '}': '{', '>': '<'}
var openMatch = map[rune]rune{'(': ')', '[': ']', '{': '}', '<': '>'}

func Day10(args []string) int {
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

	score := 0
	for i, line := range input {
		if r, ok, _ := checkCorruption(line); !ok {
			fmt.Println("FAILED LINE", i, "WITH", string(r), line)
			score += scoreRune(r)
		}
	}

	fmt.Println(score)

	return 0
}

func part2(args []string) int {
	input := tools.Read(args[0])

	scores := make([]int, 0)
	for _, line := range input {
		if _, ok, s := checkCorruption(line); ok {
			remainder := complete(s)
			scores = append(scores, scoreRemainder(remainder))
		}
	}

	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(scores[len(scores)/2])

	return 0

}

func scoreRemainder(runes []rune) int {
	acc := 0
	for _, r := range runes {
		acc *= 5
		switch r {
		case ')':
			acc += 1
		case ']':
			acc += 2
		case '}':
			acc += 3
		case '>':
			acc += 4
		}
	}

	fmt.Println(string(runes), "SCORE", acc)

	return acc
}

func complete(s stack) []rune {
	rem := make([]rune, 0)
	nStack, r, ok := s.pop()
	for ok {
		s = nStack
		if close, ok := openMatch[r]; ok {
			rem = append(rem, close)
		}

		// probably a better way of doing this...
		nStack, r, ok = s.pop()
	}

	return rem
}

func checkCorruption(input string) (rune, bool, stack) {
	s := make(stack, 0)
	for _, r := range []rune(input) {
		if open, ok := closeMatch[r]; ok {
			if nStack, curr, ok := s.pop(); ok {
				s = nStack
				if curr == open {
					continue
				} else {
					//bad close
					return r, false, s
				}
			} else {
				//empty stack
				return r, false, s
			}
		} else {
			s = s.push(r)
		}
	}

	// We could still have items in the stack, but instructions don't cover this case
	return 0, true, s
}

func scoreRune(r rune) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	default:
		return 0
	}
}

func (s stack) pop() (stack, rune, bool) {
	length := len(s)
	if length == 0 {
		return s, 0, false
	}

	result := s[length-1]
	updated := s[:length-1]

	return updated, result, true
}

func (s stack) push(item rune) stack {
	return append(s, item)
}
