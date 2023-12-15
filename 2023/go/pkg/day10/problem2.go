package day10

import (
	"fmt"
	"strings"
)

func SolveProblem2(lines []string) string {
	pipeNetwork := [][]rune{}
	startX := 0
	startY := 0

	for i, line := range lines {
		pipeNetwork = append(pipeNetwork, []rune(line))

		if idx := strings.IndexRune(line, 'S'); idx != -1 {
			startX = idx
			startY = i
		}
	}

	currentX, currentY, lastDirection := getNextElement(startX, startY, pipeNetwork)

	steps := 1

	for pipeNetwork[currentY][currentX] != 'S' {
		currentElement := pipeNetwork[currentY][currentX]

		currentX, currentY, lastDirection = walk(currentElement, currentX, currentY, lastDirection)

		steps++
	}

	return fmt.Sprint(steps / 2)
}

// !IN PROBLEM 1:

// func walk(currentElement rune, currentX, currentY int, lastDirection string) (int, int, string) {
// 	switch currentElement {
// 	case '|':
// 		if lastDirection == "up" {
// 			currentY--
// 		} else {
// 			currentY++
// 		}
// 	case '-':
// 		if lastDirection == "left" {
// 			currentX--
// 		} else {
// 			currentX++
// 		}
// 	case 'L':
// 		if lastDirection == "left" {
// 			currentY--
// 			lastDirection = "up"
// 		} else {
// 			currentX++
// 			lastDirection = "right"
// 		}
// 	case 'J':
// 		if lastDirection == "right" {
// 			currentY--
// 			lastDirection = "up"
// 		} else {
// 			currentX--
// 			lastDirection = "left"
// 		}
// 	case '7':
// 		if lastDirection == "right" {
// 			currentY++
// 			lastDirection = "down"
// 		} else {
// 			currentX--
// 			lastDirection = "left"
// 		}
// 	case 'F':
// 		if lastDirection == "left" {
// 			currentY++
// 			lastDirection = "down"
// 		} else {
// 			currentX++
// 			lastDirection = "right"
// 		}
// 	}

// 	return currentX, currentY, lastDirection
// }

// func getNextElement(startX, startY int, pipeNetwork [][]rune) (int, int, string) {
// 	el := ' '

// 	for yOffset := -1; yOffset <= 1; yOffset += 2 {
// 		el = pipeNetwork[startY+yOffset][startX]

// 		if el != '.' && isConnectedToStart(el, "vertical", 0) {
// 			dir := ""

// 			if yOffset == -1 {
// 				dir = "up"
// 			} else {
// 				dir = "down"
// 			}

// 			return startX, startY + yOffset, dir
// 		}
// 	}

// 	for xOffset := -1; xOffset <= 1; xOffset += 2 {
// 		el = pipeNetwork[startY][startX+xOffset]

// 		if el != '.' && isConnectedToStart(el, "horizontal", xOffset) {
// 			dir := ""

// 			if xOffset == -1 {
// 				dir = "left"
// 			} else {
// 				dir = "right"
// 			}

// 			return startX + xOffset, startY, dir
// 		}
// 	}

// 	return startX, startY, "none"
// }

// func isConnectedToStart(el rune, direction string, xOffset int) bool {
// 	switch direction {
// 	case "vertical":
// 		return el == '|'
// 	case "horizontal":
// 		isConnected := el == '-'

// 		if xOffset == -1 {
// 			isConnected = el == 'L' || el == 'F'
// 		}

// 		if xOffset == 1 {
// 			isConnected = el == 'J' || el == '7'
// 		}

// 		return isConnected
// 	}

// 	return false
// }
