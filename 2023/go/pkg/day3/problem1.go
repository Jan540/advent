package day3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func SolveProblem1(lines []string) string {
	sum := 0

	for i, line := range lines {
		var temp strings.Builder

		for j, r := range line {
			if unicode.IsDigit(r) {
				temp.WriteRune(r)
			}

			if (!unicode.IsDigit(r) || j == len(line)-1) && temp.Len() > 0 {
				number, err := strconv.Atoi(temp.String())

				if err != nil {
					panic(err)
				}

				if isAdjacentToSymbol(j-temp.Len(), i, temp.String(), lines) {
					sum += number
				}

				temp.Reset()
			}
		}
	}

	return fmt.Sprint(sum)
}

func isAdjacentToSymbol(x, y int, num string, lines []string) bool {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= len(num); dx++ {
			ny, nx := y+dy, x+dx
			if ny >= 0 && ny < len(lines) && nx >= 0 && nx < len(lines[ny]) && !unicode.IsDigit(rune(lines[ny][nx])) && lines[ny][nx] != '.' {
				return true
			}
		}
	}
	return false
}
