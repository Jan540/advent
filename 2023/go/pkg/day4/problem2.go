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
		newCount := processCard(line, lines)

		for _, v := range newCount {
			sum += v
		}
	}

	return fmt.Sprint(sum)
}

func processCard(card string, lines []string) map[string]int {
	cards := strings.Split(card, ": ")
	id, _ := strconv.Atoi(strings.Fields(cards[0])[1])
	allnums := strings.Split(cards[1], " | ")

	cardCount := map[string]int{
		strconv.Itoa(id): 1,
	}

	winning := strings.Fields(allnums[0])
	mynums := strings.Fields(allnums[1])

	count := 0

	for _, num := range mynums {
		if slices.Contains(winning, num) {
			count++
		}
	}

	for i := id; i < count+id; i++ {
		newCount := processCard(lines[i], lines)

		for k, v := range newCount {
			cardCount[k] += v
		}
	}

	return cardCount
}
