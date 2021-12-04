package d4

import (
	"fmt"
	"github.com/BrennanMacKay/aoc-2021/internal/tools"
	"strconv"
	"strings"
)

const boardSize = 5

type index struct {
	x int
	y int
}

type board struct {
	content     map[int]index
	unpicked    map[int]struct{}
	rowCount    [boardSize]int
	columnCount [boardSize]int
}

type winner struct {
	boardidx int
	pickIdx  int
	board    board
	pick     int
}

func Day4(args []string) int {
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
	var input = tools.Read(args[0])
	picks := readPicks(input)
	boards := readBoards(input[2:])

	fmt.Println(picks)
	fmt.Println(boards)

	winner := playGame(boards, picks)

	fmt.Println(winner)
	fmt.Println(calcScore(winner[0]))

	return 0
}

func part2(args []string) int {
	var input = tools.Read(args[0])
	picks := readPicks(input)
	boards := readBoards(input[2:])

	fmt.Println(picks)

	var winners []winner
	for {
		winners = playGame(boards, picks)
		fmt.Printf("\nWINNER idx %d pick %d pick idx %d\n\n", winners[0].boardidx, winners[0].pick, winners[0].pickIdx)
		if len(boards) == 1 {
			break
		} else {
			picks = picks[winners[0].pickIdx+1:]
			for _, v := range winners {
				delete(boards, v.boardidx)
			}
		}
	}

	fmt.Println(len(boards))
	fmt.Println(winners[0].boardidx, winners[0].pick)
	fmt.Println(calcScore(winners[0]))

	return 0
}

func calcScore(winner winner) int {
	fmt.Println(winner)
	sum := 0
	for rem := range winner.board.unpicked {
		sum += rem
	}
	return sum * winner.pick
}

func playGame(boards map[int]board, picks []int) []winner {
	fmt.Println("New Game!", picks)
	for pickIdx, pick := range picks {
		fmt.Println("Pick:", pick)
		var win = make([]winner, 0)
		for i, board := range boards {
			if index, ok := board.content[pick]; ok {
				fmt.Println("Board", i, "Has", pick)
				delete(board.unpicked, pick)
				board.columnCount[index.x] += 1
				board.rowCount[index.y] += 1
				boards[i] = board
				fmt.Println("\tColumns", board.columnCount)
				fmt.Println("\tRows", board.rowCount)
				if board.columnCount[index.x] == boardSize || board.rowCount[index.y] == boardSize {
					win = append(win, winner{i, pickIdx, board, pick})
				}
			}
		}
		if len(win) > 0 {
			fmt.Println("WINNER COUNT", len(win))
			return win
		}
	}
	panic(1)
}

func readPicks(input []string) []int {
	var picks = make([]int, 0)
	for _, v := range strings.Split(input[0], ",") {
		var parsed, _ = strconv.Atoi(v)
		picks = append(picks, parsed)
	}

	return picks
}

// requires a board to start on first line
func readBoards(input []string) map[int]board {
	var boards = make(map[int]board, 0)
	for counter := 0; len(input) > 0; counter++ {
		board := readBoard(input[0:boardSize])
		boards[counter] = board

		if len(input) <= boardSize {
			input = nil
		} else {
			input = input[boardSize+1:]
		}
	}

	return boards
}

// breaks after reading boardSize rows
func readBoard(input []string) board {
	board := board{content: make(map[int]index, boardSize*boardSize), unpicked: make(map[int]struct{}, boardSize*boardSize)}

	for i, row := range input {
		if i > boardSize {
			break
		}

		for j, item := range strings.Fields(row) {
			num, _ := strconv.Atoi(item)
			board.unpicked[num] = struct{}{}
			board.content[num] = index{x: j, y: i}
		}
	}

	return board
}
