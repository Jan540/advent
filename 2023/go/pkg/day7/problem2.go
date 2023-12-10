package day7

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

// type hand struct {
// 	cards string
// 	score int
// 	bid   int
// }

var cardValuesWithJoker = map[string]int{
	"A": 14,
	"K": 13,
	"Q": 12,
	"T": 11,
	"9": 10,
	"8": 9,
	"7": 8,
	"6": 7,
	"5": 6,
	"4": 5,
	"3": 4,
	"2": 3,
	"J": 2,
}

func SolveProblem2(lines []string) string {
	hands := []hand{}

	for _, line := range lines {
		args := strings.Split(line, " ")
		cards := args[0]
		bid, _ := strconv.Atoi(args[1])

		score := 0

		if strings.ContainsRune(cards, 'J') {
			for card := range cardValuesWithJoker {
				newScore := scoreHand(strings.ReplaceAll(cards, "J", card))

				if newScore > score {
					score = newScore
				}
			}
		} else {
			score = scoreHand(cards)
		}

		hands = append(hands, hand{
			cards,
			score,
			bid,
		})
	}

	sortHands(hands)

	totalWinnings := 0

	for i, hand := range hands {
		totalWinnings += hand.bid * (i + 1)
	}

	return fmt.Sprint(totalWinnings)
}

func scoreHand(cards string) int {
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

	return score
}

func sortHands(hands []hand) {
	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score != hands[j].score {
			return hands[i].score < hands[j].score
		}

		for k, card := range hands[i].cards {
			if card != rune(hands[j].cards[k]) {
				return cardValuesWithJoker[string(card)] < cardValuesWithJoker[string(hands[j].cards[k])]
			}
		}

		return false
	})
}
