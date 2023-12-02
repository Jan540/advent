package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveProblem2(lines []string) string {
	sum := 0

	for _, line := range lines {
		game := strings.Split(line, ": ")
		sets := strings.Split(game[1], "; ")
		maxCubes := map[string]int{}

		for _, setStr := range sets {
			cubesStr := strings.Split(setStr, ", ")

			for _, cubeStr := range cubesStr {
				cube := strings.Split(cubeStr, " ")
				name := cube[1]
				count, _ := strconv.Atoi(cube[0])

				if count > maxCubes[name] {
					maxCubes[name] = count
				}
			}
		}

		sum += maxCubes["red"] * maxCubes["green"] * maxCubes["blue"]
	}

	return fmt.Sprint(sum)
}
