package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"advent/pkg/day1"
	"advent/pkg/day2"
)

func main() {
	dayPtr := flag.Int("day", 1, "day to run")
	problemPtr := flag.Int("problem", 1, "problem to run")
	flag.Parse()

	file, err := os.ReadFile(fmt.Sprintf("/home/jan/advent/2023/go/in/day%d.in", *dayPtr))

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	result := ""

	switch *dayPtr {
	case 1:
		switch *problemPtr {
		case 1:
			fmt.Println("Solving day 1, problem 1")
			result = day1.SolveProblem1(lines)
		case 2:
			fmt.Println("Solving day 1, problem 2")
			result = day1.SolveProblem2(lines)
		default:
			result = "Invalid problem number"
		}

	case 2:
		switch *problemPtr {
		case 1:
			fmt.Println("Solving day 2, problem 1")
			result = day2.SolveProblem1(lines)
		case 2:
			fmt.Println("Solving day 2, problem 2")
			result = day2.SolveProblem2(lines)
		}
	default:
		result = "Invalid day number"
	}

	fmt.Println(result)
}
