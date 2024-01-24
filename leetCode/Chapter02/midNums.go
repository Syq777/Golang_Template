package Chapter02

/*
计算总和：首先计算数组 nums 的总和。

遍历数组：从数组的第一个元素开始遍历。

计算左右和：在遍历的过程中，我们可以维护一个变量来记录从数组开始到当前下标的所有元素的和（称为“左和”）。右侧所有元素的和可以通过总和减去左和再减去当前元素的值来计算。

比较左右和：在每一步中，检查左和是否等于右和（即总和减去左和再减去当前元素）。如果相等，那么当前下标就是我们要找的中心下标。

返回结果：如果找到了符合条件的下标，则返回它。如果遍历完整个数组后没有找到符合条件的下标，则返回 -1。
*/
func midNums(nums []int) int {
	// 先计算总和
	sum := 0
	for _, num := range nums {
		sum += num
	}
	leftSum := 0
	for i, num := range nums {
		if i != 0 {
			leftSum = nums[i-1]
		}
		rightSum := sum - num - leftSum
		if leftSum == rightSum {
			return i
		}
	}
	return -1
}
