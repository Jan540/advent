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

		for i, number := range numbers {
			firstIdx := strings.Index(line, number)
			lastIdx := strings.LastIndex(line, number)

			if firstIdx == -1 {
				continue
			}

			if digitIndices.first == "" || firstIdx < digitIndices.firstIdx {
				_, err := strconv.Atoi(number)

				if err == nil {
					digitIndices.first = number
				} else {
					digitIndices.first = fmt.Sprint(i + 1)
				}

				digitIndices.firstIdx = firstIdx
			}

			if digitIndices.last == "" || lastIdx > digitIndices.lastIdx {
				_, err := strconv.Atoi(number)

				if err == nil {
					digitIndices.last = number
				} else {
					digitIndices.last = fmt.Sprint(i + 1)
				}

				digitIndices.lastIdx = lastIdx
			}
		}

		num, _ := strconv.Atoi(string(digitIndices.first) + string(digitIndices.last))
		sum += num
	}

	return fmt.Sprint(sum)
}
