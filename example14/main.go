package main

import "fmt"

func Less(num1, num2 int) bool {
	if num1 > num2 {
		return true
	}
	return false
}

func Swap(nums *[]int, i, j int) {
	var temp int
	temp = (*nums)[i]
	(*nums)[i] = (*nums)[j]
	(*nums)[j] = temp
}

func BubbleSort(nums []int) []int {
	length := len(nums)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			if Less(nums[i], nums[j]) {
				Swap(&nums, i, j)
			}
		}
	}
	return nums
}

func main() {
	d := []int{3, 6, 2, 5, 7}
	BubbleSort(d)
	fmt.Println(d) //[1, 2, 3, 4, 6
}

//func main() {
//
//
//	srv := gin.Default()
//
//	itemCtr := controller.NewItem()
//	srv.GET("/items", itemCtr.All)
//	srv.GET("/item", itemCtr.Get)
//
//	srv.Run()
//}
