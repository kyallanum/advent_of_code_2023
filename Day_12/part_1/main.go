package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var chann = make(chan int)

func main() {
	file, _ := os.Open("input2.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	totalPossibilities := 0
	for scanner.Scan() {
		go resolveLine(scanner.Text())
		totalPossibilities += <-chann
	}
	fmt.Println(totalPossibilities)
}

func resolveLine(line string) {
	springInfo := strings.Split(line, " ")
	springHintString := strings.Split(springInfo[1], ",")
	springHint := make([]int, len(springHintString))
	for i := 0; i < len(springHint); i++ {
		springHint[i], _ = strconv.Atoi(springHintString[i])
	}

	springs := strings.Split(springInfo[0], "")
	fmt.Println("Current Line:", springs)
	chann <- getPossibilities(springs, springHint)
}

func getPossibilities(springLine []string, springHints []int) int {
	possiblePermutations := 0
	possibleIndexes := getIndexes(springLine, "?")

	for _, index := range possibleIndexes {
		springLine[index] = "."
	}

	for !allIndexesAreChar(springLine, possibleIndexes, "#") {
		for _, index := range possibleIndexes {
			if springLine[index] == "." {
				springLine[index] = "#"
				if lineMatchesHints(springLine, springHints) {
					possiblePermutations++
				}
				break
			} else if springLine[index] == "#" {
				springLine[index] = "."
			}
		}
	}
	fmt.Println(possiblePermutations)
	return possiblePermutations
}

func getIndexes(line []string, charToSearch string) []int {
	indexes := make([]int, 0)
	for index, currentChar := range line {
		if currentChar == charToSearch {
			indexes = append(indexes, index)
		}
	}
	return indexes
}

func allIndexesAreChar(line []string, indexes []int, character string) bool {
	for _, index := range indexes {
		if line[index] != character {
			return false
		}
	}
	return true
}

func lineMatchesHints(line []string, hints []int) bool {
	splitLine := make([]int, 0)
	currentSliceIndex := 0
	inGroup := false

	for _, value := range line {
		if value == "#" {
			if !inGroup {
				splitLine = append(splitLine, 1)
				inGroup = true
				continue
			} else {
				splitLine[currentSliceIndex]++
			}
		} else if value == "." {
			if inGroup {
				currentSliceIndex++
			}
			inGroup = false
		}
	}

	if len(splitLine) == len(hints) {
		for i := 0; i < len(splitLine); i++ {
			if splitLine[i] != hints[i] {
				return false
			}
		}
	} else {
		return false
	}
	fmt.Println("        Line:", line, "Hints:", hints, "Permutation:", splitLine)
	return true
}
