package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
)

var chann = make(chan int)

func main() {
	fileBytes, _ := os.ReadFile("./input.txt")
	fileContents := string(fileBytes)

	fileSlice := strings.Split(fileContents, "\n")
	totalPossibilites := 0

	wg := new(sync.WaitGroup)

	for _, line := range fileSlice {
		wg.Add(1)
		go func(line string) {
			defer wg.Done()
			resolveLine(line)
		}(line)
	}

	totalPossibilites += <-chann
	wg.Wait()

	fmt.Println(totalPossibilites)
}

func resolveLine(line string) {
	springInfo := strings.Split(line, " ")
	springHintString := strings.Split(springInfo[1], ",")
	fmt.Println(springHintString)
	springHint := make([]int, 0)
	for i := 0; i < len(springHintString); i++ {
		hint, _ := strconv.Atoi(springHintString[i])
		springHint = append(springHint, hint)
	}

	originalHint := make([]int, len(springHint))
	copy(originalHint, springHint)
	for i := 0; i < 5; i++ {
		springHint = append(springHint, originalHint...)
	}

	springs := strings.Split(springInfo[0], "")
	originalSprings := make([]string, len(springs))
	copy(originalSprings, springs)
	for i := 0; i < 5; i++ {
		springs = append(springs, append([]string{"?"}, originalSprings...)...)
	}
	fmt.Println("Current Line:", springs, "Hints:", springHint)
	chann <- getPossibilities(springs, springHint)
}

func getPossibilities(springLine []string, springHints []int) int {
	possiblePermutations := 0
	possibleIndexes := getIndexes(springLine, "?")

	for _, index := range possibleIndexes {
		springLine[index] = "."
	}

	if lineMatchesHints(springLine, springHints) {
		possiblePermutations++
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
	// fmt.Println("        Line:", line, "Hints:", hints, "Permutation:", splitLine)
	return true
}
