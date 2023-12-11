package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"advent/pkg/day1"
	"advent/pkg/day2"
	"advent/pkg/day3"
	"advent/pkg/day4"
	"advent/pkg/day5"
	"advent/pkg/day6"
	"advent/pkg/day7"
	"advent/pkg/day8"
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

	result := "Problem doesn't exist"

	fmt.Printf("Solving day %d, problem %d\n", *dayPtr, *problemPtr)
	switch *dayPtr {
	case 1:
		switch *problemPtr {
		case 1:
			result = day1.SolveProblem1(lines)
		case 2:
			result = day1.SolveProblem2(lines)
		}

	case 2:
		switch *problemPtr {
		case 1:
			result = day2.SolveProblem1(lines)
		case 2:
			result = day2.SolveProblem2(lines)
		}

	case 3:
		switch *problemPtr {
		case 1:
			result = day3.SolveProblem1(lines)
		case 2:
			result = day3.SolveProblem2(lines)
		}

	case 4:
		switch *problemPtr {
		case 1:
			result = day4.SolveProblem1(lines)
		case 2:
			result = day4.SolveProblem2(lines)
		}

	case 5:
		switch *problemPtr {
		case 1:
			result = day5.SolveProblem1(lines)
		case 2:
			result = day5.SolveProblem2(lines)
		}

	case 6:
		switch *problemPtr {
		case 1:
			result = day6.SolveProblem1(lines)
		case 2:
			result = day6.SolveProblem2(lines)
		}

	case 7:
		switch *problemPtr {
		case 1:
			result = day7.SolveProblem1(lines)
		case 2:
			result = day7.SolveProblem2(lines)
		}

	case 8:
		switch *problemPtr {
		case 1:
			result = day8.SolveProblem1(lines)
		case 2:
			result = day8.SolveProblem2(lines)
		}
	default:
		result = "Invalid day number"
	}

	fmt.Println(result)
}
