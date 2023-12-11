package main

import (
	"fmt"
	"os"
	"strings"
)

type Pipe struct {
	possibleSource [][]int
	possibleDest   [][]int
}

var possiblePipes = map[string]Pipe{
	"|": {
		possibleSource: [][]int{{-1, 0}, {1, 0}},
		possibleDest:   [][]int{{1, 0}, {-1, 0}},
	},
	"-": {
		possibleSource: [][]int{{0, -1}, {0, 1}},
		possibleDest:   [][]int{{0, 1}, {0, -1}},
	},
	"L": {
		possibleSource: [][]int{{-1, 0}, {0, 1}},
		possibleDest:   [][]int{{0, 1}, {-1, 0}},
	},
	"J": {
		possibleSource: [][]int{{-1, 0}, {0, -1}},
		possibleDest:   [][]int{{0, -1}, {-1, 0}},
	},
	"7": {
		possibleSource: [][]int{{0, -1}, {1, 0}},
		possibleDest:   [][]int{{1, 0}, {0, -1}},
	},
	"F": {
		possibleSource: [][]int{{0, 1}, {1, 0}},
		possibleDest:   [][]int{{1, 0}, {0, 1}},
	},
}

func findStarting(pipes [][]string) (int, int) {
	for row := 0; row < len(pipes); row++ {
		for column := 0; column < len(pipes[row]); column++ {
			if pipes[row][column] == "S" {
				return row, column
			}
		}
	}
	return 0, 0
}

func main() {
	fileBytes, _ := os.ReadFile("./input.txt")
	fileContents := string(fileBytes)
	pipes := make([][]string, 0)

	for _, line := range strings.Split(fileContents, "\n") {
		currentRow := make([]string, 0)
		currentRow = append(currentRow, (strings.Split(line, ""))...)
		pipes = append(pipes, currentRow)
	}

	previousRow, previousCol := findStarting(pipes)
	currentRow, currentCol := previousRow, previousCol

	possibleDirections := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for _, direction := range possibleDirections {
		row := direction[0]
		col := direction[1]
		moved := false

		possibleSources := possiblePipes[pipes[previousRow+row][previousCol+col]].possibleSource
		for _, possibleSource := range possibleSources {
			currentSource := []int{possibleSource[0] * -1, possibleSource[1] * -1}
			if currentSource[0] == direction[0] && currentSource[1] == direction[1] {
				currentRow, currentCol = currentRow+row, currentCol+col
				moved = true
				break
			}
		}

		if moved {
			break
		}
	}

	//Part 1
	totalMovements := 1
	loopCoordinates := make([][]int, 0)

	for pipes[currentRow][currentCol] != "S" {
		for _, direction := range possiblePipes[pipes[currentRow][currentCol]].possibleDest {
			destRow := direction[0]
			destCol := direction[1]

			if currentRow+destRow != previousRow || currentCol+destCol != previousCol {
				previousRow, previousCol = currentRow, currentCol
				loopCoordinates = append(loopCoordinates, []int{currentRow, currentCol})
				currentRow += destRow
				currentCol += destCol
				totalMovements++
				break
			}
		}
	}
	remainder := totalMovements % 2
	fmt.Println((totalMovements / 2) + remainder)
}
