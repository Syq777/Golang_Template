package Chapter01

// ArrayConcatenation 数组串联
func ArrayConcatenation(num []int) []int {
	nums := make([]int, len(num)*2)
	for i := 0; i < len(num); i++ {
		nums[i] = num[i]
		nums[i+len(num)] = nums[i]
	}
	return nums
}

// ArrayConcatenationPlus 优化版本
func ArrayConcatenationPlus(num []int) []int {
	DoubleLength := len(num) * 2
	return append(append(make([]int, 0, DoubleLength), num...), num...) //长度是0，容量是DoubleLength ...表示每个元素单独追加
}
