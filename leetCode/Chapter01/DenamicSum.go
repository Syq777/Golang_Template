package Chapter01

// DanamicSum 动态数组和
func DanamicSum(nums []int) []int {
	ResNums := make([]int, len(nums)) //FIXme: 这里使用切片
	ResNums[0] = nums[0]
	for i := 1; i < len(nums); i++ { //这里不适合用forrange
		ResNums[i] = nums[i] + ResNums[i-1]

	}
	return ResNums
}
