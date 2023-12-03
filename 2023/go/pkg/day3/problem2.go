package day3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type gear struct {
	x, y int
}

func SolveProblem2(lines []string) string {
	sum := 0
	gears := map[gear]([]int){}

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

				adjacent, gear := isAdjacentToGear(j-temp.Len(), i, temp.String(), lines)

				if adjacent {
					gears[gear] = append(gears[gear], number)
				}

				temp.Reset()
			}
		}
	}

	for _, gear := range gears {
		if len(gear) == 2 {
			sum += gear[0] * gear[1]
		}
	}

	return fmt.Sprint(sum)
}

func isAdjacentToGear(x, y int, num string, lines []string) (bool, gear) {
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= len(num); dx++ {
			ny, nx := y+dy, x+dx
			if ny >= 0 && ny < len(lines) &&
				nx >= 0 && nx < len(lines[ny]) &&
				lines[ny][nx] == '*' {
				return true, gear{nx, ny}
			}
		}
	}
	return false, gear{}
}
