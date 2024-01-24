package Chapter02

// PlusOne
/*函数 addOneToNumber(数组 digits):
  从数组的末尾开始遍历数组
      如果当前元素小于 9:
          将该元素加一
          返回数组
      否则:
          将该元素设置为 0

  所有元素都被设置为 0 后:
  在数组的开头插入一个 1
  返回数组
*/
// 要实现的是只有最后一位加一
func PlusOne(nums []int) []int {
	// 如果当前位是9，当前位置0，并传下一位+1
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < 9 {
			nums[i]++
			return nums
		}
		nums[i] = 0
	}
	//全部都为9的情况
	// 检查nums[0]
	if nums[0] == 0 {
		newNums := make([]int, len(nums)+1)
		newNums[0] = 1
		return newNums
	}
	return nums
}
