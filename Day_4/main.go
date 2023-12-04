package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Card struct {
	CardID         int
	CardNumbers    []int
	WinningNumbers []int
}

// func (card *Card) getCardPoints() int {
// 	currentNumbers := card.CardNumbers
// 	winningNumbers := card.WinningNumbers

// 	var numWinningNumbers int

// 	for _, number := range currentNumbers {

// 		if slices.Contains(winningNumbers, number) {
// 			numWinningNumbers++
// 		}
// 	}
// 	// fmt.Println(card.CardNumbers)
// 	// fmt.Println(card.WinningNumbers)

// 	return int(math.Pow(2.0, float64(numWinningNumbers-1)))
// }

func (card *Card) getNumWinningNumbers() int {
	currentNumbers := card.CardNumbers
	winningNumbers := card.WinningNumbers

	var numWinningNumbers int
	for _, number := range currentNumbers {
		if slices.Contains(winningNumbers, number) {
			numWinningNumbers++
		}
	}

	// fmt.Println("ID:", card.CardID, "Winning Numbers:", numWinningNumbers)
	return numWinningNumbers
}

func main() {
	file, _ := os.Open("./input.txt")
	scanner := bufio.NewScanner(file)
	defer file.Close()

	// Part 1
	// var totalPoints int
	// for scanner.Scan() {
	// 	currentText := scanner.Text()
	// 	// fmt.Println(currentText)
	// 	newCard := parseCard(currentText)
	// 	currentPoints := newCard.getCardPoints()
	// 	// fmt.Println(currentPoints)
	// 	totalPoints += currentPoints
	// }

	// Part 2
	cardMap := make(map[int]*Card)
	for scanner.Scan() {
		currentText := scanner.Text()
		// fmt.Println(currentText)
		newCard := parseCard(currentText)
		cardMap[newCard.CardID] = newCard
	}

	fmt.Println(processCards(cardMap))
}

func parseCard(line string) *Card {
	cardInfo := strings.Split(line, ":")
	spaceIndex := strings.LastIndex(cardInfo[0], " ")
	cardID, _ := strconv.Atoi(cardInfo[0][spaceIndex+1:])

	re := regexp.MustCompile(`\s+`)
	cardInfo[1] = re.ReplaceAllString(cardInfo[1], " ")
	cardInfo[1] = cardInfo[1] + " "

	cardNumbers := strings.Split(cardInfo[1], "|")
	cardCurrentNumbers := strings.Split(cardNumbers[0][1:len(cardNumbers[0])-1], " ")
	cardWinningNumbers := strings.Split(cardNumbers[1][1:len(cardNumbers[1])-1], " ")

	cardCurrentNumbersInt := make([]int, 0)
	cardWinningNumbersInt := make([]int, 0)

	for _, cardNumber := range cardCurrentNumbers {
		currentNumber, _ := strconv.Atoi(cardNumber)
		cardCurrentNumbersInt = append(cardCurrentNumbersInt, currentNumber)
	}

	for _, cardNumber := range cardWinningNumbers {
		winningNumber, _ := strconv.Atoi(cardNumber)
		cardWinningNumbersInt = append(cardWinningNumbersInt, winningNumber)
	}

	// fmt.Println("Current Numbers:", cardCurrentNumbersInt, "WinningNumbers:", cardWinningNumbersInt)

	return &Card{
		CardID:         cardID,
		CardNumbers:    cardCurrentNumbersInt,
		WinningNumbers: cardWinningNumbersInt,
	}
}

func processCards(cards map[int]*Card) int {
	cardAmounts := make(map[int][]int)
	for cardID := range cards {
		// Entry = numCards, numWinningNumbers
		cardAmounts[cardID] = make([]int, 2)
		cardAmounts[cardID] = []int{1, 0}
	}

	cardIDs := maps.Keys(cards)
	slices.Sort(cardIDs)

	for _, ID := range cardIDs {
		// fmt.Println("ID:", ID, "Card:", cardAmounts[ID])
		cardAmounts[ID][1] = cards[ID].getNumWinningNumbers()
		// fmt.Println("Current Card: ", cardAmounts[ID])
		for i := 1; i <= cardAmounts[ID][1] && i < len(cardAmounts); i++ {
			cardAmounts[ID+i][0] += cardAmounts[ID][0]
			// fmt.Println("Card:", ID+i, "Value:", cardAmounts[ID+i])
		}
	}

	var totalCards int
	for _, cardData := range cardAmounts {
		totalCards += cardData[0]
	}
	return totalCards
}
