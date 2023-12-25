package Chapter01

func TotalAssets(nums [][]int) int {
	maxWealth := 0

	for _, num := range nums {
		wealth := 0
		for _, nu := range num { // FIxme：遍历两次
			wealth += nu
		}
		if wealth > maxWealth {
			maxWealth = wealth
		}
	}
	return maxWealth
}
