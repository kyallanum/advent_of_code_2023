package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Hand struct {
	Cards []string
	Bet   int
	Type  int
}

func (hand *Hand) HandType() int {
	//Return 7: 5 of a kind
	//Return 6: 4 of a kind
	//Return 5: Full House
	//Return 4: 3 of a kind
	//Return 3: 2 pair
	//Return 2: 1 pair
	//return 1: High card

	dict := make(map[string]int)
	for _, num := range hand.Cards {
		dict[num] = dict[num] + 1
	}

	numJs := dict["J"]
	if numJs == 5 {
		return 7
	}
	delete(dict, "J")

	keys := maps.Keys(dict)
	maxMapKey := keys[0]

	for _, key := range keys {
		if dict[key] > dict[maxMapKey] {
			maxMapKey = key
		}
	}

	dict[maxMapKey] += numJs

	if len(keys) == 1 {
		hand.Type = 7
		return 7
	} else if len(keys) == 2 {
		if dict[keys[0]] == 4 || dict[keys[1]] == 4 {
			hand.Type = 6
			return 6
		} else if (dict[keys[0]] == 3 && dict[keys[1]] == 2) || (dict[keys[0]] == 2 && dict[keys[1]] == 3) {
			hand.Type = 5
			return 5
		}
	} else if len(keys) == 3 {
		if dict[keys[0]] == 3 || dict[keys[1]] == 3 || dict[keys[2]] == 3 {
			hand.Type = 4
			return 4
		} else if (dict[keys[0]] == 2 && dict[keys[1]] == 2) || (dict[keys[0]] == 2 && dict[keys[2]] == 2) || (dict[keys[1]] == 2 && dict[keys[2]] == 2) {
			hand.Type = 3
			return 3
		}
	} else if len(keys) == 4 {
		hand.Type = 2
		return 2
	}
	hand.Type = 1
	return 1
}

func main() {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	hands := make([]Hand, 0)
	for scanner.Scan() {
		currentHand := parseLine(scanner.Text())
		hands = append(hands, currentHand)
	}

	// fmt.Println(hands)

	sort.Slice(hands, func(i, j int) bool {
		handIType := hands[i].HandType()
		handJType := hands[j].HandType()
		cardMap := map[string]string{
			"A": "14",
			"K": "13",
			"Q": "12",
			"T": "10",
			"J": "1",
		}
		if handIType == handJType {
			for index := 0; index < 5; index++ {
				currentCardIString, success := cardMap[hands[i].Cards[index]]
				if !success {
					currentCardIString = hands[i].Cards[index]
				}
				currentCardI, _ := strconv.Atoi(currentCardIString)

				currentCardJString, success := cardMap[hands[j].Cards[index]]
				if !success {
					currentCardJString = hands[j].Cards[index]
				}
				currentCardJ, _ := strconv.Atoi(currentCardJString)

				if currentCardI == currentCardJ {
					continue
				} else if currentCardI < currentCardJ {
					return true
				}
				return false
			}
		}
		return handIType < handJType
	})

	fmt.Println(hands)

	totalScore := 0
	for i := len(hands) - 1; i >= 0; i-- {
		totalScore += hands[i].Bet * (i + 1)
	}

	fmt.Println(totalScore)
}

func parseLine(line string) Hand {
	lineContentsSlice := strings.Split(line, " ")
	currentHand := Hand{
		Cards: make([]string, 0),
		Bet:   0,
	}

	currentBet, _ := strconv.Atoi(lineContentsSlice[1])
	currentHand.Bet = currentBet

	currentCards := strings.Split(lineContentsSlice[0], "")
	currentHand.Cards = append(currentHand.Cards, currentCards...)

	return currentHand
}
