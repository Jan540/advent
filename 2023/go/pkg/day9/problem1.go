package day9

import (
	"fmt"
	"strconv"
	"strings"
)

func SolveProblem1(lines []string) string {
	sum := 0

	for _, line := range lines {
		numsStr := strings.Split(line, " ")
		nums := make([]int, len(numsStr))

		for i, numStr := range numsStr {
			nums[i], _ = strconv.Atoi(numStr)
		}

		tree := descent(nums)
		sum += ascend(tree)
	}

	return fmt.Sprint(sum)
}

func descent(nums []int) [][]int {
	result := [][]int{nums}
	newRow := []int{}

	for i, num := range nums {
		if i == len(nums)-1 {
			break
		}

		newRow = append(newRow, nums[i+1]-num)
	}

	if allZeros(newRow) {
		return append(result, newRow)
	}

	newRows := descent(newRow)
	return append(result, newRows...)
}

func ascend(tree [][]int) int {
	current := 0

	for i := len(tree) - 1; i >= 0; i-- {
		current += tree[i][len(tree[i])-1]
	}

	return current
}

func allZeros(newRow []int) bool {
	for _, num := range newRow {
		if num != 0 {
			return false
		}
	}

	return true
}
