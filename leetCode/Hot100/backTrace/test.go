package main

import (
	"fmt"
)

// permute generates all permutations of the given slice.
func permute(nums []int) [][]int {
	var result [][]int
	backtrack(&result, nums, 0)
	return result
}

// backtrack is the helper function for generating permutations.
func backtrack(result *[][]int, nums []int, start int) {
	if start == len(nums) {
		tmp := make([]int, len(nums))
		copy(tmp, nums)
		*result = append(*result, tmp)
		return
	}
	for i := start; i < len(nums); i++ {
		nums[start], nums[i] = nums[i], nums[start]
		backtrack(result, nums, start+1)
		nums[start], nums[i] = nums[i], nums[start] // backtrack
	}
}

func main() {
	nums := []int{1, 2, 3}
	permutations := permute(nums)
	fmt.Println(permutations)
}
