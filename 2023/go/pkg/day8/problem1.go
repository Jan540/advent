package day8

import (
	"fmt"
	"strings"
)

type node struct {
	left  string
	right string
}

func SolveProblem1(lines []string) string {
	locations := map[string]node{}
	instructions := strings.Split(lines[0], "")

	for _, line := range lines[2:] {
		args := strings.Split(line, " = ")
		code := args[0]
		clean_dirs := strings.Replace(args[1], "(", "", -1)
		cleaner_dirs := strings.Replace(clean_dirs, ")", "", -1)
		directions := strings.Split(cleaner_dirs, ", ")

		locations[code] = node{directions[0], directions[1]}
	}

	currentNode := locations["AAA"]
	steps := 0

	for i := 0; i < len(instructions); i++ {
		instruction := instructions[i]
		dir := ""

		if instruction == "L" {
			dir = currentNode.left
		} else if instruction == "R" {
			dir = currentNode.right
		} else {
			panic("Invalid instruction")
		}

		steps++

		if dir == "ZZZ" {
			break
		}

		currentNode = locations[dir]

		if i == len(instructions)-1 {
			i = -1
		}
	}

	return fmt.Sprint(steps)
}
