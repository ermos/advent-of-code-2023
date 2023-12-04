package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

type Card struct {
	Matches  []string
	Instance int
}

func main() {
	var results int
	var cards []*Card

	for _, s := range utils.LoadInput("4") {
		card := Card{
			Instance: 1,
		}

		sets := strings.Split(strings.Split(s, ":")[1], "|")
		winningNumbers := strings.Fields(sets[0])
		myNumbers := strings.Fields(sets[1])

		for _, winningNumber := range winningNumbers {
			for _, myNumber := range myNumbers {
				if winningNumber == myNumber {
					card.Matches = append(card.Matches, myNumber)
				}
			}
		}

		cards = append(cards, &card)
	}

	for index, card := range cards {
		if len(card.Matches) > 0 {
			for i := index + 1; i < index+1+len(card.Matches); i++ {
				cards[i].Instance += card.Instance
			}
		}
	}

	for _, card := range cards {
		results += card.Instance
	}

	fmt.Println(results)
}
