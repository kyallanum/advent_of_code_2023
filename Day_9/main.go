package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	total := 0
	for scanner.Scan() {
		total += resolveLine(scanner.Text())
	}

	fmt.Println(total)
}

func allZeros(numArray []int) bool {
	for _, number := range numArray {
		if number != 0 {
			return false
		}
	}
	return true
}

// Part 1
// func resolveLine(line string) int {
// 	entries := make([][]int, 0)
// 	numbersStringArray := strings.Split(line, " ")
// 	numbers := make([]int, len(numbersStringArray))

// 	entries = append(entries, numbers)

// 	for index, numString := range numbersStringArray {
// 		numbers[index], _ = strconv.Atoi(numString)
// 	}

// 	for !allZeros(entries[len(entries)-1]) {
// 		newEntry := make([]int, len(entries[len(entries)-1])-1)
// 		lastEntry := entries[len(entries)-1]

// 		for i, j := 0, 1; j < len(lastEntry); i, j = i+1, j+1 {
// 			newEntry[i] = lastEntry[j] - lastEntry[i]
// 		}

// 		entries = append(entries, newEntry)
// 	}

// 	entries[len(entries)-1] = append(entries[len(entries)-1], 0)

// 	for i := len(entries) - 2; i >= 0; i-- {
// 		lowerRowLastNumber := entries[i+1][len(entries[i+1])-1]
// 		lastNumber := entries[i][len(entries[i])-1]
// 		entries[i] = append(entries[i], lastNumber+lowerRowLastNumber)
// 	}

// 	return entries[0][len(entries[0])-1]
// }

func resolveLine(line string) int {
	entries := make([][]int, 0)
	numbersStringArray := strings.Split(line, " ")
	numbers := make([]int, len(numbersStringArray))

	entries = append(entries, numbers)

	for index, numString := range numbersStringArray {
		numbers[index], _ = strconv.Atoi(numString)
	}

	for !allZeros(entries[len(entries)-1]) {
		newEntry := make([]int, len(entries[len(entries)-1])-1)
		lastEntry := entries[len(entries)-1]

		for i, j := 0, 1; j < len(lastEntry); i, j = i+1, j+1 {
			newEntry[i] = lastEntry[j] - lastEntry[i]
		}

		entries = append(entries, newEntry)
	}

	entries[len(entries)-1] = append([]int{0}, entries[len(entries)-1]...)
	for i := len(entries) - 2; i >= 0; i-- {
		firstNumber := entries[i][0]
		firstNumberLowerRow := entries[i+1][0]
		entries[i] = append([]int{firstNumber - firstNumberLowerRow}, entries[i]...)
	}
	fmt.Println(entries)

	return entries[0][0]
}
