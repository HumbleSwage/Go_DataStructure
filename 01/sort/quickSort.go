package main

import "fmt"

// 快速排序
/*
1.left 表示数组左边的下标
2.right 表示数组右边的下标
3.array 表示要排序的数组
*/
func QuickSort(left int, right int, array *[6]int) {
	l := left
	r := right
	//pivot是中轴(支点)的意思
	pivot := array[(left+right)/2]
	temp := 0

	//将比pivot小的数放到左边，将比pivot大的数放在右边
	for l < r {
		//从pivot的左边找到大于等于pivot的值
		for array[l] < pivot {
			l++
		}
		//从pivot的右边找到小于等于pivot的值
		for array[r] > pivot {
			r--
		}
		//表明本次任务结束
		if l >= r {
			break
		}
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	//如果l==rr再移动一位
	if l == r {
		l++
		r--
	}
	//向左递归
	if left < r {
		QuickSort(left, r, array)
	}
	//向右递归
	if right > l {
		QuickSort(l, right, array)
	}
}

func main() {

	arr := [6]int{-9, 78, 23, -567, 70}
	fmt.Println("初始化的数组：", arr)
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println("排序后的数组：", arr)
}
