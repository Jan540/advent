package day1

import (
	"fmt"
	"strconv"
	"unicode"
)

func SolveProblem1(lines []string) string {
	sum := 0

	for _, line := range lines {
		var first, last rune

		for _, char := range line {
			if !unicode.IsDigit(char) {
				continue
			}

			if first == 0 {
				first = char
			}

			last = char
		}

		num, _ := strconv.Atoi(string(first) + string(last))
		sum += num
	}

	return fmt.Sprint(sum)
}
