package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Pipe struct {
	possibleDest [][]int
}

var possiblePipes = map[string]Pipe{
	"|": {
		possibleDest: [][]int{{1, 0}, {-1, 0}},
	},
	"-": {
		possibleDest: [][]int{{0, 1}, {0, -1}},
	},
	"L": {
		possibleDest: [][]int{{0, 1}, {-1, 0}},
	},
	"J": {
		possibleDest: [][]int{{0, -1}, {-1, 0}},
	},
	"7": {
		possibleDest: [][]int{{1, 0}, {0, -1}},
	},
	"F": {
		possibleDest: [][]int{{1, 0}, {0, 1}},
	},
}

type Tile struct {
	Value   string
	Visited bool
}

func findStarting(tiles *[][]Tile) (int, int) {
	for row := 0; row < len(*tiles); row++ {
		for column := 0; column < len((*tiles)[row]); column++ {
			if (*tiles)[row][column].Value == "S" {
				return row, column
			}
		}
	}
	return 0, 0
}

func parseToTiles(fileContents string) [][]Tile {
	rows := strings.Split(fileContents, "\n")

	tiles := make([][]Tile, 0)

	for _, row := range rows {
		cols := strings.Split(row, "")
		rowTiles := make([]Tile, 0)
		for _, col := range cols {
			newTile := &Tile{
				Value:   col,
				Visited: false,
			}
			rowTiles = append(rowTiles, *newTile)
		}
		tiles = append(tiles, rowTiles)
	}

	return tiles
}

func firstMove(startingRow int, startingCol int, tiles *[][]Tile) (int, int) {
	possibleDirections := map[string][][]int{
		"|": {{-1, 0}, {1, 0}},
		"-": {{0, -1}, {0, 1}},
		"L": {{1, 0}, {0, -1}},
		"J": {{1, 0}, {0, -1}},
		"7": {{-1, 0}, {0, 1}},
		"F": {{-1, 0}, {0, 1}},
	}

	directions := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	for _, direction := range directions {
		directionRow := direction[0]
		directionCol := direction[1]

		if startingRow+directionRow < 0 || startingRow+directionRow >= len((*tiles)[startingRow]) || startingCol+directionCol < 0 || startingCol+directionCol >= len((*tiles)[startingRow]) {
			continue
		}

		newTile := (*tiles)[startingRow+directionRow][startingCol+directionCol].Value

		for _, direction := range possibleDirections[newTile] {
			if directionRow == direction[0] && directionCol == direction[1] {
				startingRow += directionRow
				startingCol += directionCol
				(*tiles)[startingRow][startingCol].Visited = true
				return startingRow, startingCol
			}
		}
	}
	return 0, 0
}

func runLoop(currentRow, currentCol, startingRow, startingCol int, tiles *[][]Tile) {
	numMovements := 1
	for (*tiles)[currentRow][currentCol].Value != "S" {
		currentPipe := (*tiles)[currentRow][currentCol].Value
		directions := possiblePipes[currentPipe]
		for _, direction := range directions.possibleDest {
			newRow := direction[0]
			newCol := direction[1]
			if (currentRow+newRow != startingRow || currentCol+newCol != startingCol) && !(*tiles)[currentRow+newRow][currentCol+newCol].Visited {
				startingRow = currentRow
				startingCol = currentCol
				currentRow += newRow
				currentCol += newCol
				numMovements++
				(*tiles)[currentRow][currentCol].Visited = true
				break
			}
		}
	}
}

func translateUnusedPipeToGround(tiles *[][]Tile) {
	for rowIndex := range *tiles {
		for colIndex := range (*tiles)[rowIndex] {
			if !(*tiles)[rowIndex][colIndex].Visited {
				(*tiles)[rowIndex][colIndex].Value = "."
			}
		}
	}
}

func replaceS(startingRow, startingCol int, tiles *[][]Tile) {
	accepts := map[string][][]int{
		"|": {{-1, 0}, {1, 0}},
		"-": {{0, -1}, {0, 1}},
		"L": {{1, 0}, {0, -1}},
		"J": {{1, 0}, {0, 1}},
		"7": {{-1, 0}, {0, 1}},
		"F": {{-1, 0}, {0, -1}},
	}

	destinations := map[string][][]int{
		"|": {{1, 0}, {-1, 0}},
		"-": {{0, 1}, {0, -1}},
		"L": {{0, 1}, {-1, 0}},
		"J": {{0, -1}, {-1, 0}},
		"7": {{1, 0}, {0, -1}},
		"F": {{1, 0}, {0, 1}},
	}

	possibleDirections := [][]int{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1},
	}

	startingDirections := make([][]int, 0)

	for _, direction := range possibleDirections {
		directionRow := direction[0]
		directionCol := direction[1]

		if startingRow+directionRow < 0 ||
			startingRow+directionRow >= len(*tiles) ||
			startingCol+directionCol < 0 ||
			startingCol+directionCol >= len((*tiles)[startingRow]) {
			continue
		}

		newKey := (*tiles)[startingRow+directionRow][startingCol+directionCol]
		values := accepts[newKey.Value]
		for _, value := range values {
			valueRow := value[0]
			valueColumn := value[1]

			if directionRow == valueRow && directionCol == valueColumn {
				newDirection := []int{directionRow, directionCol}
				startingDirections = append(startingDirections, newDirection)
				break
			}
		}
	}

	for key := range destinations {
		compareDestinations := destinations[key]
		if reflect.DeepEqual(compareDestinations, startingDirections) || reflect.DeepEqual(compareDestinations, [][]int{startingDirections[1], startingDirections[0]}) {
			(*tiles)[startingRow][startingCol].Value = key
			break
		}
	}
}

func getEnclosedTiles(tiles *[][]Tile) int {
	totalEnclosed := 0
	for _, row := range *tiles {
		enclosed := false
		previousPipe := "-"
		for i := 0; i < len(row); i++ {
			if row[i].Value == "|" {
				enclosed = !enclosed
				continue
			}
			if row[i].Value == "-" {
				continue
			}
			if row[i].Value == "." {
				if enclosed {
					totalEnclosed++
				}
				continue
			}
			if previousPipe == "-" && (row[i].Value == "J" || row[i].Value == "L" || row[i].Value == "7" || row[i].Value == "F") {
				previousPipe = row[i].Value
				continue
			} else {
				if previousPipe == "L" && row[i].Value == "7" {
					enclosed = !enclosed
				}

				if previousPipe == "F" && row[i].Value == "J" {
					enclosed = !enclosed
				}
				previousPipe = "-"
				continue
			}
		}
	}

	return totalEnclosed
}

func main() {
	fileBytes, _ := os.ReadFile("./input.txt")
	fileContents := string(fileBytes)

	boardTiles := parseToTiles(fileContents)

	startingRow, startingCol := findStarting(&boardTiles)
	firstMoveRow, firstMoveCol := firstMove(startingRow, startingCol, &boardTiles)

	runLoop(firstMoveRow, firstMoveCol, startingRow, startingCol, &boardTiles)
	translateUnusedPipeToGround(&boardTiles)
	for _, row := range boardTiles {
		for _, col := range row {
			fmt.Printf("%s ", col.Value)
		}
		fmt.Print("\n")
	}
	replaceS(startingRow, startingCol, &boardTiles)

	totalEnclosed := getEnclosedTiles(&boardTiles)

	fmt.Println("Num Enclosed:", totalEnclosed)

	fmt.Println("Finished")
}
