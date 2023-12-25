package main

import "fmt"

// 声明切片的方式
func testSlice() {
	// 直接用make
	slice := make([]int, 3)
	fmt.Println(slice)
	// 直接声明
	slice2 := []int{1, 2, 3}
	fmt.Println(slice2)
	// 从数组或另一个切片创建
	arr := []int{1, 2, 3}
	slice3 := arr[1:2] //从1开始打印第二个元素
	fmt.Println(slice3)

}
func main() {
	testSlice()
}
