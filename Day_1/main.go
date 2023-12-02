package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// //Day 1! Let's goooooo
	var sumCalibrationValues int
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		calibrationValue := getFirstLastDigits(scanner.Text())
		// fmt.Println("String:", scanner.Text(), "Value:", calibrationValue)
		sumCalibrationValues += calibrationValue

	}
	fmt.Println(sumCalibrationValues)
}

func getFirstLastDigits(input string) int {
	numberMap := map[string]int{
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
	}

	firstMatchIndex := len(input)
	lastMatchIndex := -1
	var firstMatch string
	var lastMatch string

	for key := range numberMap {
		currentFirstIndex := strings.Index(input, key)
		currentLastIndex := strings.LastIndex(input, key)
		if currentFirstIndex != -1 {
			if currentFirstIndex < firstMatchIndex {
				firstMatchIndex = currentFirstIndex
				firstMatch = key

			}
			if currentLastIndex > lastMatchIndex {
				lastMatchIndex = currentLastIndex
				lastMatch = key
			}
		}
	}

	return 10*numberMap[firstMatch] + numberMap[lastMatch]
}
