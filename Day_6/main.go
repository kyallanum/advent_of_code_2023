package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fileBytes, _ := os.ReadFile("./input.txt")
	fileContents := string(fileBytes)
	times, distances := parseFile(fileContents)

	totalTime := 1

	for index, time := range times {
		for i := 0; i < time/2; i++ {
			// fmt.Println("(time-i)*i:", (time-i)*i, "distance:", distances[index])
			if (time-i)*i > distances[index] {
				totalTime *= time - i*2 + 1
				// fmt.Println(time - (i * 2) + 1)
				break
			}
		}
	}

	fmt.Println(totalTime)
}

func parseFile(fileContents string) ([]int, []int) {
	lines := strings.Split(fileContents, "\n")
	re := regexp.MustCompile(`\s+`)
	fmt.Println(fileContents)
	for index, line := range lines {
		lines[index] = re.ReplaceAllString(line, "")
		fmt.Println(lines[index])
	}

	timeString := strings.Split(strings.Split(lines[0], ":")[1], " ")
	distanceString := strings.Split(strings.Split(lines[1], ":")[1], " ")
	fmt.Println(timeString)
	fmt.Println(distanceString)

	time := make([]int, len(timeString))
	distance := make([]int, len(distanceString))

	for index, _ := range timeString {
		time[index], _ = strconv.Atoi(timeString[index])
		distance[index], _ = strconv.Atoi(distanceString[index])
	}

	return time, distance
}
