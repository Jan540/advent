package day5

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func SolveProblem1(lines []string) string {
	seeds := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	locations := make([]int, len(seeds))

	// convert seeds to []int
	for _, seedStr := range seeds {
		seed, err := strconv.Atoi(seedStr)

		if err != nil {
			panic(err)
		}

		num := seed

		for i, line := range lines[2:] {

			if line == "" || unicode.IsDigit(rune(line[0])) {
				lines = lines[i+2:]
				continue
			}

			args := strings.Fields(line)
			destStart, _ := strconv.Atoi(args[0])
			srcStart, _ := strconv.Atoi(args[1])
			rng, _ := strconv.Atoi(args[2])

			if num >= srcStart && num < srcStart+rng {
				num = destStart + (num - srcStart)
			}
		}

		locations = append(locations, num)
	}

	return fmt.Sprint(locations)
}
