package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Game struct {
	GameId int
	Reds   []int
	Greens []int
	Blues  []int
}

func (game *Game) getMaxRed() int {
	reds := game.Reds
	sort.Ints(reds)

	return reds[len(reds)-1]
}

func (game *Game) getMaxGreen() int {
	greens := game.Greens
	sort.Ints(greens)

	return greens[len(greens)-1]
}

func (game *Game) getMaxBlue() int {
	blues := game.Blues
	sort.Ints(blues)

	return blues[len(blues)-1]
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Part 1
	// maxColors := []int{12, 13, 14}
	// var sumGameIds int
	// for scanner.Scan() {
	// 	currentGame := parseGame(scanner.Text())
	// 	if currentGame.getMaxRed() <= maxColors[0] && currentGame.getMaxGreen() <= maxColors[1] && currentGame.getMaxBlue() <= maxColors[2] {
	// 		sumGameIds += currentGame.GameId
	// 	}
	// }
	// fmt.Println(sumGameIds)

	//Part 2
	var sumPowers int
	for scanner.Scan() {
		currentGame := parseGame(scanner.Text())
		currentPower := currentGame.getMaxRed() * currentGame.getMaxGreen() * currentGame.getMaxBlue()
		sumPowers += currentPower
	}

	fmt.Println(sumPowers)
}

func parseGame(line string) *Game {
	currentGame := &Game{
		Reds:   make([]int, 0),
		Greens: make([]int, 0),
		Blues:  make([]int, 0),
	}

	gameData := strings.Split(line, ":")
	gameID, _ := strconv.Atoi(gameData[0][5:])
	// fmt.Println(gameData)
	currentGame.GameId = gameID

	turns := strings.Split(gameData[1], ";")
	// fmt.Println(turns)
	for _, turn := range turns {
		colors := strings.Split(turn, ",")
		// fmt.Println(colors)
		for _, color := range colors {
			colorData := strings.Split(color, " ")
			colorData = colorData[1:]
			// fmt.Println(colorData)
			number, _ := strconv.Atoi(colorData[0])
			switch colorData[1] {
			case "red":
				currentGame.Reds = append(currentGame.Reds, number)
			case "green":
				currentGame.Greens = append(currentGame.Greens, number)
			case "blue":
				currentGame.Blues = append(currentGame.Blues, number)
			}
		}
		// fmt.Println(currentGame)
	}

	return currentGame
}
