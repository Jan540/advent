package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveProblem1(lines []string) string {
	sum := 0

	for _, line := range lines {
		game := strings.Split(line, ": ")
		id, _ := strconv.Atoi(strings.Split(game[0], " ")[1])
		sets := strings.Split(game[1], "; ")
		valid := true

		for _, setStr := range sets {
			cubesStr := strings.Split(setStr, ", ")
			cubes := map[string]int{}

			for _, cubeStr := range cubesStr {
				cube := strings.Split(cubeStr, " ")
				name := cube[1]
				count, _ := strconv.Atoi(cube[0])
				cubes[name] = count
			}

			if cubes["red"] > 12 || cubes["green"] > 13 || cubes["blue"] > 14 {
				valid = false
				break
			}
		}

		if valid {
			sum += id
		}
	}

	return fmt.Sprint(sum)
}
