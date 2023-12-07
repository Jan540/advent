package day6

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveProblem2(lines []string) string {
	times := strings.Fields(strings.Split(lines[0], ": ")[1])
	distances := strings.Fields(strings.Split(lines[1], ": ")[1])

	wins := 0

	time, _ := strconv.Atoi(strings.Join(times, ""))
	distanceToBeat, _ := strconv.Atoi(strings.Join(distances, ""))

	for speed := 0; speed <= time; speed++ {
		distance := speed * (time - speed)

		if distance > distanceToBeat {
			wins = time - (speed * 2) + 1
			break
		}
	}

	return fmt.Sprint(wins)
}
