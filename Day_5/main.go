package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Almanac struct {
	SeedToSoil        [][]int
	SoilToFertilizer  [][]int
	FertilizerToWater [][]int
	WaterToLight      [][]int
	LightToTemp       [][]int
	TempToHumidity    [][]int
	HumidityToLoc     [][]int
}

func (alm *Almanac) getSeedLocation(seed int) int {
	// fmt.Println("Seed:", seed)
	var soil int
	soilFound := false
	for _, entry := range alm.SeedToSoil {
		if entry[1] <= seed && seed < entry[1]+entry[2] {
			diff := seed - entry[1]
			soil = entry[0] + diff
			soilFound = true
			break
		}
	}
	if !soilFound {
		soil = seed
	}

	// fmt.Println("Soil:", soil)
	var fertilizer int
	fertilizerFound := false
	for _, entry := range alm.SoilToFertilizer {
		if entry[1] <= soil && soil < entry[1]+entry[2] {
			diff := soil - entry[1]
			fertilizer = entry[0] + diff
			fertilizerFound = true
		}
	}
	if !fertilizerFound {
		fertilizer = soil
	}

	// fmt.Println("Fertilizer:", fertilizer)
	var water int
	waterFound := false
	for _, entry := range alm.FertilizerToWater {
		if entry[1] <= fertilizer && fertilizer < entry[1]+entry[2] {
			diff := fertilizer - entry[1]
			water = entry[0] + diff
			waterFound = true
		}
	}
	if !waterFound {
		water = fertilizer
	}

	// fmt.Println("Water:", water)
	var light int
	lightFound := false
	for _, entry := range alm.WaterToLight {
		if entry[1] <= water && water < entry[1]+entry[2] {
			diff := water - entry[1]
			light = entry[0] + diff
			lightFound = true
		}
	}
	if !lightFound {
		light = water
	}

	// fmt.Println("Light:", light)
	var temp int
	tempFound := false
	for _, entry := range alm.LightToTemp {
		if entry[1] <= light && light < entry[1]+entry[2] {
			diff := light - entry[1]
			temp = entry[0] + diff
			tempFound = true
		}
	}
	if !tempFound {
		temp = light
	}

	// fmt.Println("Temp:", temp)
	var humidity int
	humidityFound := false
	for _, entry := range alm.TempToHumidity {
		if entry[1] <= temp && temp < entry[1]+entry[2] {
			diff := temp - entry[1]
			humidity = entry[0] + diff
			humidityFound = true
		}
	}
	if !humidityFound {
		humidity = temp
	}

	// fmt.Println("Humid:", humidity)
	var location int
	locationFound := false
	for _, entry := range alm.HumidityToLoc {
		if entry[1] <= humidity && humidity < entry[1]+entry[2] {
			diff := humidity - entry[1]
			location = entry[0] + diff
			locationFound = true
		}
	}
	if !locationFound {
		location = humidity
	}

	// fmt.Println("Loc:", location)
	return location
}

var chann = make(chan int)

func main() {
	fileBytes, _ := os.ReadFile("./input.txt")
	fileContents := string(fileBytes)
	seeds, almanac := parseAlmanac(fileContents)

	closestSeed := math.MaxInt32

	// Part 1
	// for _, seed := range seeds {
	// 	closestSeed = int(math.Min(float64(closestSeed), float64(almanac.getSeedLocation(seed))))
	// }

	//Part 2
	for i := 0; i < len(seeds); i += 2 {
		go processSeed(i, seeds, almanac)
		seedLoc := <-chann
		if seedLoc < closestSeed {
			closestSeed = seedLoc
		}
	}

	fmt.Println(closestSeed)
}

func processSeed(seedIndex int, seeds []int, almanac *Almanac) {
	lowestLoc := math.MaxInt32

	for i := 0; i < seeds[seedIndex+1]; i++ {
		currentLoc := almanac.getSeedLocation(seeds[seedIndex] + i)
		if currentLoc < lowestLoc {
			lowestLoc = currentLoc
		}
	}

	chann <- lowestLoc
}

func parseAlmanac(fileString string) ([]int, *Almanac) {
	fileContents := strings.Split(fileString, "\n")
	seedInfo := strings.Split(fileContents[0], ":")
	seedsString := strings.Split(seedInfo[1][1:], " ")
	seeds := make([]int, len(seedsString))

	for i, seedString := range seedsString {
		currentSeed, _ := strconv.Atoi(seedString)
		seeds[i] = int(currentSeed)
	}

	almanac := &Almanac{
		SeedToSoil:        make([][]int, 0),
		SoilToFertilizer:  make([][]int, 0),
		FertilizerToWater: make([][]int, 0),
		WaterToLight:      make([][]int, 0),
		LightToTemp:       make([][]int, 0),
		TempToHumidity:    make([][]int, 0),
		HumidityToLoc:     make([][]int, 0),
	}

	newSection := true
	mapType := ""
	for i := 1; i < len(fileContents); i++ {
		currentLine := fileContents[i]

		if currentLine == "" {
			newSection = true
			continue
		}
		if newSection {
			mapType = strings.Split(currentLine, " ")[0]
			newSection = false
			continue
		}

		lineSlice := strings.Split(currentLine, " ")
		// fmt.Println(lineSlice)
		destStart, _ := strconv.Atoi(lineSlice[0])
		sourceStart, _ := strconv.Atoi(lineSlice[1])
		numRange, _ := strconv.Atoi(lineSlice[2])

		switch mapType {
		case "seed-to-soil":
			almanac.SeedToSoil = append(almanac.SeedToSoil, []int{destStart, sourceStart, numRange})
		case "soil-to-fertilizer":
			almanac.SoilToFertilizer = append(almanac.SoilToFertilizer, []int{destStart, sourceStart, numRange})
		case "fertilizer-to-water":
			almanac.FertilizerToWater = append(almanac.FertilizerToWater, []int{destStart, sourceStart, numRange})
		case "water-to-light":
			almanac.WaterToLight = append(almanac.WaterToLight, []int{destStart, sourceStart, numRange})
		case "light-to-temperature":
			almanac.LightToTemp = append(almanac.LightToTemp, []int{destStart, sourceStart, numRange})
		case "temperature-to-humidity":
			almanac.TempToHumidity = append(almanac.TempToHumidity, []int{destStart, sourceStart, numRange})
		case "humidity-to-location":
			almanac.HumidityToLoc = append(almanac.HumidityToLoc, []int{destStart, sourceStart, numRange})
		}
	}
	return seeds, almanac
}
