package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveProblem2(lines []string) string {
	sum := 0

	for _, line := range lines {
		numsStr := strings.Split(line, " ")
		nums := make([]int, len(numsStr))

		for i, numStr := range numsStr {
			nums[i], _ = strconv.Atoi(numStr)
		}

		tree := descent(nums)
		sum += ascendBackwards(tree)
	}

	return fmt.Sprint(sum)
}

func ascendBackwards(tree [][]int) int {
	current := 0

	for i := len(tree) - 1; i >= 0; i-- {
		current = tree[i][0] - current
	}

	return current
}

// IN PROBLEM 1:
// func descent(nums []int) [][]int {
// 	result := [][]int{nums}
// 	newRow := []int{}

// 	for i, num := range nums {
// 		if i == len(nums)-1 {
// 			break
// 		}

// 		newRow = append(newRow, nums[i+1]-num)
// 	}

// 	if allZeros(newRow) {
// 		return append(result, newRow)
// 	}

// 	newRows := descent(newRow)
// 	return append(result, newRows...)
// }

// func allZeros(newRow []int) bool {
// 	for _, num := range newRow {
// 		if num != 0 {
// 			return false
// 		}
// 	}

// 	return true
// }
