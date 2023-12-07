package day6

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveProblem1(lines []string) string {
	times := strings.Fields(strings.Split(lines[0], ": ")[1])
	distances := strings.Fields(strings.Split(lines[1], ": ")[1])

	wins := 1

	for i, timeStr := range times {
		time, _ := strconv.Atoi(timeStr)
		distanceToBeat, _ := strconv.Atoi(distances[i])

		count := 0

		for speed := 0; speed <= time; speed++ {
			distance := speed * (time - speed)

			fmt.Println(speed, distance, distanceToBeat)

			if distance > distanceToBeat {
				count++
			}
		}

		wins *= count
	}

	return fmt.Sprint(wins)
}
