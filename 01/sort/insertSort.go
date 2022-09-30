package main

import (
	"fmt"
)

func InsertSort(arr *[7]int) {
	//注意这里是从1开始的
	for i := 1; i < len(arr); i++ {
		//完成第i次，给第二个元素找到合适位置并插入
		insertVal := arr[i]
		//insertIndex始终是需要调整位置元素的前一个元素的位置
		insertIndex := i - 1
		//从大到小
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			//数据后移，不用担心覆盖，因为我们已经将其值用insertVal进行了保存
			arr[insertIndex+1] = arr[insertIndex]
			insertIndex--
		}
		//插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		//打印一下
		fmt.Printf("第%d次插入后%v\n", i, *arr)
	}
}

func main() {
	fmt.Println("实现插入排序")
	arr := [7]int{23, 0, 13, 56, 34, 45, 9}
	InsertSort(&arr)
}
