package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/exp/maps"
)

type Map struct {
	desertMap  map[string][]string
	directions []string
}

func (desertMap *Map) processMap() int {

	directionIndex := 0

	totalKeys := maps.Keys(desertMap.desertMap)
	startingKeys := make([]string, 0)

	for _, key := range totalKeys {
		if string(key[2]) == "A" {
			startingKeys = append(startingKeys, key)
		}
	}

	numIterationsPerKey := make([]int, len(startingKeys))

	for index := range startingKeys {
		numIterations := 0
		for {
			if directionIndex == len(desertMap.directions) {
				directionIndex = 0
			}

			if string(startingKeys[index][2]) == "Z" {
				break
			}

			switch desertMap.directions[directionIndex] {
			case "R":
				startingKeys[index] = desertMap.desertMap[startingKeys[index]][1]
			case "L":
				startingKeys[index] = desertMap.desertMap[startingKeys[index]][0]
			}

			directionIndex++
			numIterations++
		}
		numIterationsPerKey[index] = numIterations
	}

	return getLeastCommonMultiple(numIterationsPerKey)
}

func getLeastCommonMultiple(numbers []int) int {
	lcm := numbers[0]
	for i := 0; i < len(numbers); i++ {
		num1 := lcm
		num2 := numbers[i]
		gcd := 1
		for num2 != 0 {
			temp := num2
			num2 = num1 % num2
			num1 = temp
		}
		gcd = num1
		lcm = (lcm * numbers[i]) / gcd
	}

	return lcm
}

func main() {
	fileBytes, _ := os.ReadFile("./input.txt")
	fileContents := string(fileBytes)
	instructions, desertMap := parseFile(fileContents)

	desertMapInstructions := &Map{
		desertMap:  desertMap,
		directions: instructions,
	}

	fmt.Println(desertMapInstructions.processMap())
}

func parseFile(fileContents string) ([]string, map[string][]string) {
	fileInfo := strings.Split(fileContents, "\n\n")
	instructionsStrings := strings.Split(fileInfo[0], "")

	desertMapStrings := strings.Split(fileInfo[1], "\n")
	var desertMap = map[string][]string{}

	for _, line := range desertMapStrings {
		key := strings.TrimSpace(strings.Split(line, "=")[0])
		valueString := strings.TrimSpace(strings.Split(line, "=")[1])
		valueSlice := strings.Split(valueString[1:len(valueString)-1], ",")
		valueSlice[1] = strings.TrimSpace(valueSlice[1])

		desertMap[key] = make([]string, 2)
		desertMap[key] = valueSlice
	}

	return instructionsStrings, desertMap
}
