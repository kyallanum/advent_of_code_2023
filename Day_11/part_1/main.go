package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sky := make([][]string, 0)
	for scanner.Scan() {
		currentLine := scanner.Text()
		separatedLine := strings.Split(currentLine, "")
		sky = append(sky, separatedLine)
	}

	expand(&sky)
	galaxyMap := getAllGalaxies(&sky)

	for _, line := range sky {
		for _, value := range line {
			fmt.Printf("%s", value)
		}
		fmt.Println()
	}

	fmt.Println(galaxyMap)

	totalDistances := getAllLengths(&galaxyMap)
	fmt.Println("Total Distances:", totalDistances)

}

func getAllGalaxies(sky *[][]string) map[int][]int {
	galaxyMap := make(map[int][]int)
	galaxyNumber := 1
	for rowIndex := 0; rowIndex < len(*sky); rowIndex++ {
		for colIndex := 0; colIndex < len((*sky)[rowIndex]); colIndex++ {
			if (*sky)[rowIndex][colIndex] == "#" {
				newGalaxy := []int{rowIndex, colIndex}
				galaxyMap[galaxyNumber] = newGalaxy
				galaxyNumber++
			}
		}
	}

	return galaxyMap
}

func expand(sky *[][]string) {
	for rowIndex := 0; rowIndex < len(*sky); rowIndex++ {
		if !slices.Contains((*sky)[rowIndex], "#") {
			newRow := make([]string, len((*sky)[rowIndex]))
			for index := range newRow {
				newRow[index] = "."
			}
			*sky = append((*sky)[:rowIndex], append([][]string{newRow}, (*sky)[rowIndex:]...)...) // It's fancy dancy for insert()
			rowIndex++
		}
	}

	for colIndex := 0; colIndex < len((*sky)[0]); colIndex++ {
		columnExpanded := false
		currentCol := make([]string, len(*sky))
		for i := 0; i < len(currentCol); i++ {
			currentCol[i] = (*sky)[i][colIndex]
		}

		if !slices.Contains(currentCol, "#") {
			for i := 0; i < len(*sky); i++ {
				(*sky)[i] = append((*sky)[i][:colIndex], append([]string{"."}, (*sky)[i][colIndex:]...)...) // Here we go again
				columnExpanded = true
			}
		}

		if columnExpanded {
			colIndex++
		}
	}
}

func getAllLengths(galaxyMap *map[int][]int) int {
	totalDistance := 0
	keys := make([]int, 0)
	for k := range *galaxyMap {
		keys = append(keys, k)
	}

	slices.Sort(keys)

	for i := 0; i < len(keys)-1; i++ {
		for j := i + 1; j < len(keys); j++ {
			deltaRow := ((*galaxyMap)[keys[j]][0] + 1) - ((*galaxyMap)[keys[i]][0] + 1)
			deltaColumn := ((*galaxyMap)[keys[j]][1] + 1) - ((*galaxyMap)[keys[i]][1] + 1)
			distance := int(math.Abs(float64(deltaColumn)) + math.Abs(float64(deltaRow)))
			fmt.Println("Comparison:", (*galaxyMap)[keys[i]], (*galaxyMap)[keys[j]], deltaRow, deltaColumn, "Distance:", distance)
			totalDistance += distance
		}
	}

	return totalDistance
}
