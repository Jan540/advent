package day7

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type hand struct {
	cards string
	score int
	bid   int
}

var cardValues = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 10,
	"J": 11,
	"9": 9,
	"8": 8,
	"7": 7,
	"6": 6,
	"5": 5,
	"4": 4,
	"3": 3,
	"2": 2,
}

func SolveProblem1(lines []string) string {
	hands := []hand{}

	for _, line := range lines {
		args := strings.Split(line, " ")
		cards := args[0]
		bid, _ := strconv.Atoi(args[1])
		score := 0
		searched := []rune{}

		for _, card := range cards {
			if slices.Contains(searched, card) {
				continue
			}

			searched = append(searched, card)

			count := strings.Count(cards, string(card))

			if count == 5 {
				score += 6
				break
			}

			if count == 4 {
				score += 5
				break
			}

			if count == 3 {
				score += 3
			}

			if count == 2 {
				score += 1
			}
		}

		hands = append(hands, hand{cards, score, bid})
	}

	// pls don't judge me
	// it sorts ascending
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score != hands[j].score {
			return hands[i].score < hands[j].score
		}

		for k, card := range hands[i].cards {
			if card != rune(hands[j].cards[k]) {
				return cardValues[string(card)] < cardValues[string(hands[j].cards[k])]
			}
		}

		return false
	})

	totalWinnings := 0

	for i, hand := range hands {
		totalWinnings += hand.bid * (i + 1)
	}

	return fmt.Sprint(totalWinnings)
}
