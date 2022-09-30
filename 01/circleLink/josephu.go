package main

import (
	"fmt"
)

// 创建一个小孩的结构体
type Boy struct {
	No   int
	next *Boy //指向下一个小孩的指针
}

// 编写一个函数，构成单向环形链表
// num：表示小孩的个数
// *Boy：返回该环形链表的头指针
func AddBoy(num int) *Boy {
	first := &Boy{}  //空节点
	curBoy := &Boy{} //空节点 【辅助帮忙的】
	if num < 1 {
		return first
	}
	//for循环开始构建
	for i := 0; i <= num; i++ {
		boy := &Boy{
			No: i,
			//没有设置next就为nil
		}
		//构成一个环形链表，必须要给一个辅助指针
		//1.因为第一个小孩比较特殊
		if i == 1 { //第一个小孩
			first = boy
			curBoy = boy
			curBoy.next = first //这句话的目的就是让其形成环形
		} else {
			curBoy.next = boy
			curBoy = boy
			curBoy.next = first //构成环形链表
		}
	}
	return first
}

// 显示单向环形链表
func ShowBoy(first *Boy) {
	//处理一哈环形链表为空的情况
	if first.next == nil {
		fmt.Println("链表为空,没有什么可展示的")
		return
	}
	//创建一个辅助指针来帮助遍历
	curBoy := first
	for {
		fmt.Printf(" 小孩编号=%d ==>", curBoy.No)
		//退出条件
		if curBoy.next == first {
			break
		}
		//移动curboy
		curBoy = curBoy.next
	}
}

func GetNum(first *Boy) int {
	//处理一哈环形链表为空的情况
	if first.next == nil {
		fmt.Println("链表为空,没有什么可展示的")
		return 0
	}
	//创建一个辅助指针来帮助遍历
	curBoy := first
	count := 0
	for {
		count++
		//退出条件
		if curBoy.next == first {
			break
		}
		//移动curboy
		curBoy = curBoy.next
	}
	return count
}

/*
设编号为1，2，3,...n的n个人围坐一圈，约定编号为k（1<=k<=n）的人从1开始报数，数到m
的那个人出列，它的下一位又从1开始报数，数到m的那个人又出列，一次类推，知道所有人出列
由此产生一个出队编号序列的问题.[环形列表的最后只剩下一个人]
*/
func PlayGamesfirst(first *Boy, startNo int, countNo int) {
	//1.空链表单独处理
	if first.next == nil {
		fmt.Println("环形链表为空")
		return
	}
	//判断一下startNo是不是比我们最大的编号还大
	if startNo > GetNum(first) {
		fmt.Println("起点的编号大于环形链表中的最大编号")
	}
	//2.定义辅助指针，帮助我们删除指针
	tail := first
	//3.让tail指向环形链表的最后一个小孩【这个非常重要，在删除小孩时需要用到】
	for {
		if tail.next == first { //说明tail已经到最后一个小孩了
			break
		}
		tail = tail.next
	}
	//4.让first移动到startNo【后面我删除小孩就以first为准】
	for i := 1; i <= startNo-1; i++ {
		first = first.next
		tail = tail.next
	}
	//5.开始数countNum下，然后就删除first指向的小孩【这个地方一定要清楚为什么需要一个双层循环】
	for {
		//开始数countNum-1次
		for i := 1; i <= countNo-1; i++ {
			first = first.next
			tail = tail.next
		}
		fmt.Printf("%d号小孩出圈\n", first.No)
		//删除first指向的节点【tail始终在first的后面】
		first = first.next
		tail.next = first
		//判断什么时候退出【圈中只有一个小孩】
		if first == tail {
			fmt.Printf("%d号小孩出圈\n", first.No)
			break
		}
	}
}

func main() {
	fmt.Println("用环形链表来解决约瑟夫问题")
	first := AddBoy(5)
	ShowBoy(first)
	fmt.Println()
	PlayGamesfirst(first, 2, 3)
}
