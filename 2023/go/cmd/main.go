package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"advent/pkg/day1"
)

func main() {
	dayPtr := flag.Int("day", 1, "day to run")
	problemPtr := flag.Int("problem", 1, "problem to run")
	flag.Parse()

	file, err := os.ReadFile(fmt.Sprintf("./in/day%d.in", *dayPtr))

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
	default:
		result = "Invalid day number"
	}

	fmt.Println(result)
}
