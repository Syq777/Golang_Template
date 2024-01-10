package Chapter02

// PlusOne这里是每位加1
func PlusOne(nums []int) []int {
	res := make([]int, len(nums))
	for i, num := range nums {
		res[i] = num + 1
	}
	return res
}

// 要实现的是只有最后一位加一
func PlusOne(nums []int) []int {
	res := make([]int, len(nums))
	// 如果当前位是9，当前位置0，并传下一位+1
	length := len(nums)
	for i, num := range nums {
		if nums[length-1] == 9:
			nums[length - 1] = 0
			nums
	}
	return res
}
