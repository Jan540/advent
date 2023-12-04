package day4

import (
	"fmt"
	"slices"
	"strings"
)

func SolveProblem1(lines []string) string {
	sum := 0

	for _, line := range lines {
		cards := strings.Split(line, ": ")
		allnums := strings.Split(cards[1], " | ")

		winning := strings.Fields(allnums[0])
		mynums := strings.Fields(allnums[1])

		points := 0

		for _, num := range mynums {
			if slices.Contains(winning, num) {
				if points == 0 {
					points = 1
				} else {
					points *= 2
				}
			}
		}

		sum += points
	}

	return fmt.Sprint(sum)
}
