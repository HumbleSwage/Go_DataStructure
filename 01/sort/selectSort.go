package main

import (
	"fmt"
)

func SelectSort(arr *[5]int) {
	//注意这个地方的标准的访问元素的方式是(*arr)[1] = 600
	//但是go语言的底层封住了，所以使用arr[1]也可以实现相同的效果
	//外层循环依次遍历每一个需要找到的位置
	for i := 0; i < len(arr)-1; i++ {
		max := arr[i]
		index := i
		for j := i; j < len(arr); j++ {
			//内层循环找到最大的那一个
			if arr[j] > max {
				max = arr[j]
				index = j
			} else {
				continue
			}
		}
		//交换两个元素的位置
		// arr[i] = arr[i] + arr[index]
		// arr[index] = arr[i] - arr[index]
		// arr[i] = arr[i] - arr[index]
		//go语言交换元素有一个特别方便的形式
		arr[i], arr[index] = arr[index], arr[i]
	}
}

func main() {
	fmt.Println("实现选择排序[从大到小]")
	arr := [5]int{10, 34, 19, 100, 80}
	fmt.Println(arr)

	SelectSort(&arr)
	fmt.Println(arr)
}
