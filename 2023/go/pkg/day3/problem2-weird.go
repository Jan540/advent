package day3

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func SolveProblem2Weird(lines []string) string {
	sum := 0

	for i, line := range lines {
		for j, r := range line {
			if r == '*' {
				partnums := getAdjacentNumbers(j, i, lines)

				if len(partnums) != 2 {
					continue
				}

				sum += partnums[0] * partnums[1]
			}
		}
	}

	return fmt.Sprint(sum)
}

func getAdjacentNumbers(x, y int, lines []string) []int {
	nums := []int{}

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			ny, nx := y+dy, x+dx
			if ny >= 0 && ny < len(lines) && nx >= 0 && nx < len(lines[ny]) && unicode.IsDigit(rune(lines[ny][nx])) {
				var temp strings.Builder
				for _, r := range lines[ny][nx:] {
					if unicode.IsDigit(r) {
						temp.WriteRune(r)
					} else {
						break
					}
				}

				number, _ := strconv.Atoi(temp.String())
				nums = append(nums, number)
			}
		}
	}

	fmt.Println(nums)

	return nums
}

// in day 1:
// func isAdjacentToSymbol(x, y int, num string, lines []string) bool {
// 	for dy := -1; dy <= 1; dy++ {
// 		for dx := -1; dx <= len(num); dx++ {
// 			ny, nx := y+dy, x+dx
// 			if ny >= 0 && ny < len(lines) && nx >= 0 && nx < len(lines[ny]) && !unicode.IsDigit(rune(lines[ny][nx])) && lines[ny][nx] != '.' {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }
