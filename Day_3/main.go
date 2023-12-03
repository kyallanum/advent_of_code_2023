package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	fileBytes, _ := os.ReadFile("./input.txt")
	fileContents := string(fileBytes)
	arrayFileContents := parseFile(fileContents)

	var sumNumbers int

	//An array describing each movement in directions.
	directions := [][]int{
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1}}

	// We loop through every single element in 2-dimensional array
	for yIndex, line := range arrayFileContents {
		for xIndex, char := range line {
			// Its part two, so we just need to find gears
			if char == "*" {
				// fmt.Println("Char:", char)
				// We need to keep track of each number that we find.
				foundNumbers := make([]int, 0)
				for _, direction := range directions {
					newXIndex := xIndex + direction[1]
					newYIndex := yIndex + direction[0]
					// Back up to where the number starts and then go to where number ends. Parse that window into one number. Add to "foundNumbers" if not already added.
					if unicode.IsDigit(rune(arrayFileContents[newYIndex][newXIndex][0])) {
						// fmt.Println("X:", newXIndex, "Y:", newYIndex)
						for newXIndex-1 >= 0 && unicode.IsDigit(rune(arrayFileContents[newYIndex][newXIndex-1][0])) {
							newXIndex--
						}
						numberStartIndex := newXIndex

						for newXIndex < len(line) && unicode.IsDigit(rune(arrayFileContents[newYIndex][newXIndex][0])) {
							newXIndex++
						}
						numberEndIndex := newXIndex
						newNumber, _ := strconv.Atoi(strings.Join(arrayFileContents[newYIndex][numberStartIndex:numberEndIndex], ""))
						// fmt.Println("New Number:", newNumber)
						if !slices.Contains(foundNumbers, newNumber) {
							foundNumbers = append(foundNumbers, newNumber)
						}
					}
				}
				// If it is a gear with two numbers, then get the gear ratio and add
				if len(foundNumbers) == 2 {
					gearRatio := foundNumbers[0] * foundNumbers[1]
					sumNumbers += gearRatio
				}
			}
		}
	}

	fmt.Println(sumNumbers)
}

func parseFile(fileContents string) [][]string {
	lines := strings.Split(fileContents, "\n")
	fileContentsArray := make([][]string, len(lines))

	for index, line := range lines {
		fileContentsArray[index] = strings.Split(line, "")
	}

	return fileContentsArray
}
