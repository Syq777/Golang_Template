package Chapter02

/*
太好了！让我们来看一下修正后的伪代码，这将帮助您理解整个数组旋转的过程。

```go
package Chapter02

// 主函数来旋转数组

	func rotateArray(nums []int, k int) {
	    n := len(nums)
	    // 处理 k 大于数组长度的情况
	    k = k % n

	    // 先反转整个数组
	    reverseSlice(nums, 0, n-1)

	    // 反转前 k 个元素
	    reverseSlice(nums, 0, k-1)

	    // 反转剩余的元素
	    reverseSlice(nums, k, n-1)
	}

// 反转 nums 中从 start 到 end 的部分

	func reverseSlice(nums []int, start int, end int) {
	    for start < end {
	        // 交换 nums[start] 和 nums[end]
	        nums[start], nums[end] = nums[end], nums[start]
	        start++
	        end--
	    }
	}

```

### 关键点解释：

1. **处理 `k` 的值**：
  - 如果 `k` 大于数组长度，使用 `k % n` 来确保我们不会尝试访问数组外的元素。

2. **反转整个数组**：
  - 这是将元素移动到接近它们最终位置的第一步。

3. **分别反转两部分**：
  - 首先反转数组的前 `k` 个元素。
  - 然后反转从第 `k` 个元素开始到数组结尾的部分。

4. **`reverseSlice` 函数**：
  - 这个函数接受数组和要反转的部分的起始和结束索引。
  - 通过交换元素来反转指定的数组部分。

您可以根据这个伪代码来调整您的 Go 代码。这应该能够帮助您实现正确的数组旋转功能。尝试一下并告诉我结果如何！
*/
func rotateArray(nums []int, k int) []int {
	n := len(nums)
	k = k % n
	reverseSlice(nums, 0, k-1)
	reverseSlice(nums, k, n-1)
	reverseSlice(nums, 0, n-1)
	return nums
}

func reverseSlice(nums []int, start int, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
