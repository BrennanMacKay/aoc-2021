package tools

import (
	"bufio"
	"os"
	"strconv"
)

func Read(path string) []string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func ConvertToInt(input []string) []int {
	var output = make([]int, len(input))

	for i, v := range input {
		var parsed, err = strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		output[i] = parsed
	}

	return output
}
