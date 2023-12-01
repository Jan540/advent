package day1

import (
	"fmt"
	"strconv"
	"strings"
)

type DigitIndices struct {
	first, last       string
	firstIdx, lastIdx int
}

func SolveProblem2(lines []string) string {
	sum := 0
	numbers := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	for _, line := range lines {
		var digitIndices DigitIndices

		for _, number := range numbers {
			firstIdx := strings.Index(line, number)
			lastIdx := strings.LastIndex(line, number)

			if firstIdx == -1 {
				continue
			}

			if digitIndices.first == "" || firstIdx < digitIndices.firstIdx {
				digitIndices.first = getNumber(number)
				digitIndices.firstIdx = firstIdx
			}

			if digitIndices.last == "" || lastIdx > digitIndices.lastIdx {
				digitIndices.last = getNumber(number)
				digitIndices.lastIdx = lastIdx
			}
		}

		num, _ := strconv.Atoi(string(digitIndices.first) + string(digitIndices.last))
		sum += num
	}

	return fmt.Sprint(sum)
}

func getNumber(number string) string {
	switch number {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return number
	}
}
