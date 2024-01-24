package Day1

/*
给你一个 m 行 n 列的矩阵 matrix ，请按照 顺时针螺旋顺序 ，返回矩阵中的所有元素。
*/
/*
想法：
第一行，最后一列，然后再倒着输出最后一行，依次遍历直到长度相同
*/
func spiralMatrix(matrix [][]int) []int {
	var res []int
	m := len(matrix)
	n := len(matrix[0])
	//	初始化变量
	top := 0
	left := 0
	bottom := m - 1
	right := n - 1
	//遍历
	for top <= bottom && left <= right { // 边界
		// 向右遍历
		for col := left; col <= right; col++ {
			//添加到res
			res = append(res, matrix[top][col])
		}
		top += 1
		// 向下遍历
		for row := top; row <= bottom; row++ {
			res = append(res, matrix[row][right])
		}
		right -= 1
		if top <= bottom && left <= right {
			//向左遍历
			for col := right; col >= left; col-- {
				res = append(res, matrix[bottom][col])
			}
			bottom -= 1
			// 向上遍历
			for row := bottom; row >= top; row-- {
				res = append(res, matrix[row][left])
			}
			left += 1
			//
		}
	}
	return res

}
