package day5

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func SolveProblem1(lines []string) string {
	seeds := strings.Split(strings.Split(lines[0], ": ")[1], " ")
	lowest := math.MaxInt64

	for _, seedStr := range seeds {
		seed, err := strconv.Atoi(seedStr)

		if err != nil {
			panic(err)
		}

		num := seed
		found := false

		for _, line := range lines[3:] {

			if line == "" || !unicode.IsDigit(rune(line[0])) {
				found = false
				continue
			}

			if found {
				continue
			}

			args := strings.Fields(line)

			destStart, _ := strconv.Atoi(args[0])
			srcStart, _ := strconv.Atoi(args[1])
			rng, _ := strconv.Atoi(args[2])

			if num >= srcStart && num < srcStart+rng {
				num = destStart + (num - srcStart)
				found = true
			}
		}

		if num < lowest {
			lowest = num
		}
	}

	return fmt.Sprint(lowest)
}
