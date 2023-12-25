package Chapter01

// Calculategems 输入：jewels = "aA", stones = "aAAbbbb"
// 输出：3
// Calculategems 石头和宝石
func Calculategems(jewels string, stones string) int {
	// TODO：遗漏的一点 jewels中可能会有重复字段 使用map去重
	jewelsSet := make(map[rune]bool)
	count := 0
	for _, jewel := range jewels {
		jewelsSet[jewel] = true
	}
	// 遍历stones中存在几个jewels
	for _, stone := range stones {
		if jewelsSet[stone] {
			count++
		}
	}
	return count
}
