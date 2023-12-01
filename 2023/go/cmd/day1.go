package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	file, err := os.ReadFile("./in/day1.in")

	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(file), "\n")

	numStrings := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	sum := 0

	for _, line := range lines {
		var first, last rune

		for _, char := range line {
			if !unicode.IsDigit(char) {
				continue
			}

			if first == 0 {
				first = char
			}

			last = char
		}

		num, _ := strconv.Atoi(string(first) + string(last))
		sum += num
	}

	fmt.Println(sum)
}
