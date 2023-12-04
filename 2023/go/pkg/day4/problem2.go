package day4

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func SolveProblem2(lines []string) string {
	sum := 0

	for _, line := range lines {
		sum += processCard(line, lines)
	}

	return fmt.Sprint(sum)
}

func processCard(card string, lines []string) int {
	cards := strings.Split(card, ": ")
	id, _ := strconv.Atoi(strings.Fields(cards[0])[1])
	allnums := strings.Split(cards[1], " | ")

	cardCount := 1

	winning := strings.Fields(allnums[0])
	mynums := strings.Fields(allnums[1])

	count := id

	for _, num := range mynums {
		if slices.Contains(winning, num) {
			cardCount += processCard(lines[count], lines)
			count++
		}
	}

	return cardCount
}
