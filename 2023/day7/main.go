package main

import (
	"bufio"
	"cmp"
	"fmt"
	"log"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Hand struct {
	Cards string
	Bid   int
}

func countRunes(s string) map[rune]int {
	counter := make(map[rune]int)
	for _, label := range s {
		if count, _ := counter[label]; true {
			counter[label] = count + 1
		}
	}
	return counter
}

func getCardsValue(h Hand, labelToValue map[rune]int) int {
	/*
		Treat h.Cards as a number in a numeric system with base len(labelToValue)
	*/
	value := 0
	mostSignificantIdx := len(h.Cards) - 1
	base := len(labelToValue)
	for i, label := range h.Cards {
		cardValue, _ := labelToValue[label]
		value += cardValue * int(math.Pow(float64(base), float64(mostSignificantIdx-i)))
	}
	return value
}

func replaceJokers(cards string, defaultReplacement rune) string {
	/*
		Find if there are 'J's among the cards and replace them with the best option
	*/

	if cards == "JJJJJ" {
		return strings.ReplaceAll(cards, "J", string(defaultReplacement))
	}

	// Find the most used label
	counter := countRunes(cards)
	max := -1
	var labelReplacement rune
	for label, count := range counter {
		if label != 'J' && count > max {
			labelReplacement = label
			max = count
		}
	}

	if labelReplacement != 0 {
		return strings.ReplaceAll(cards, "J", string(labelReplacement))
	}

	return cards
}

func getHandValue(h Hand, labels []rune, shouldReplaceJokers bool) int {
	labelToValue := make(map[rune]int)
	for i, l := range labels {
		labelToValue[l] = i
	}

	cardsValue := getCardsValue(h, labelToValue)

	cards := h.Cards
	if shouldReplaceJokers {
		cards = replaceJokers(cards, labels[1])
	}
	cardCounter := countRunes(cards)
	lenCardCounter := len(cardCounter)

	var handType int
	if lenCardCounter == 1 {
		handType = 6
	} else if lenCardCounter == 2 {
		for _, value := range cardCounter {
			if value == 1 || value == 4 {
				handType = 5
				break
			} else {
				handType = 4
				break
			}
		}
	} else if lenCardCounter == 3 {
		for _, value := range cardCounter {
			if value == 3 {
				handType = 3
				break
			} else if value == 2 {
				handType = 2
				break
			}
		}
	} else if lenCardCounter == 4 {
		handType = 1
	}

	handTypeValue := labelToValue[labels[handType]]
	// Add handTypeValue as the most significant digit
	return cardsValue + handTypeValue*int(math.Pow(float64(len(labels)), float64(len(h.Cards))))
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var hands []Hand

	for scanner.Scan() {
		line := scanner.Text()
		cardsAndBid := strings.Split(line, " ")
		cards, bidStr := cardsAndBid[0], cardsAndBid[1]
		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			log.Fatal(err)
		}
		hands = append(hands, Hand{Cards: cards, Bid: bid})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	cardLabels := [13]rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
	slices.SortFunc(hands, func(a, b Hand) int {
		return cmp.Compare(getHandValue(a, cardLabels[:], false), getHandValue(b, cardLabels[:], false))
	})

	s1 := 0
	for i, h := range hands {
		s1 += h.Bid * (i + 1)
	}
	fmt.Println("s1:", s1)

	cardLabelsWithJoker := [13]rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}
	slices.SortFunc(hands, func(a, b Hand) int {
		return cmp.Compare(getHandValue(a, cardLabelsWithJoker[:], true), getHandValue(b, cardLabelsWithJoker[:], true))
	})

	s2 := 0
	for i, h := range hands {
		s2 += h.Bid * (i + 1)
	}
	fmt.Println("s2:", s2)

	// fmt.Println(replaceJokers("23456", '3')) // "23456"
	// fmt.Println(replaceJokers("KTJJT", '3')) // "KTTTT"
	// fmt.Println(replaceJokers("2345J", '3')) // "23455"
	// fmt.Println(replaceJokers("JJJJJ", '2')) // "22222"
	// fmt.Println(replaceJokers("JJJJJ", '3')) // "33333"
}
