package day8

import (
	"fmt"
	"strings"
)

// in problem1.go:
// type node struct {
// 	left  string
// 	right string
// }

func SolveProblem2(lines []string) string {
	locations := map[string]node{}
	instructions := strings.Split(lines[0], "")
	starterNodes := []node{}

	for _, line := range lines[2:] {
		args := strings.Split(line, " = ")
		code := args[0]
		clean_dirs := strings.Replace(args[1], "(", "", -1)
		cleaner_dirs := strings.Replace(clean_dirs, ")", "", -1)
		directions := strings.Split(cleaner_dirs, ", ")

		locations[code] = node{directions[0], directions[1]}

		if code[len(code)-1:] == "A" {
			starterNodes = append(starterNodes, locations[code])
		}
	}

	steps := make([]int, len(starterNodes))

	for i, node := range starterNodes {
		currentNode := node

		for j := 0; j < len(instructions); j++ {
			instruction := instructions[j]
			dir := ""

			if instruction == "L" {
				dir = currentNode.left
			} else if instruction == "R" {
				dir = currentNode.right
			} else {
				panic("Invalid instruction")
			}

			steps[i]++

			if dir[len(dir)-1] == 'Z' {
				break
			}

			currentNode = locations[dir]

			if j == len(instructions)-1 {
				j = -1
			}
		}
	}

	return fmt.Sprint(findLCM(steps))
}

func findLCM(steps []int) int {
	result := steps[0]
	for i := 1; i < len(steps); i++ {
		result = lcm(result, steps[i])
	}
	return result
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
