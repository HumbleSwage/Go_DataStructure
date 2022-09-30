package main

import (
	"fmt"
)

// 定义一个猫猫结构体
type CatNode struct {
	no   int
	name string
	next *CatNode
}

// 插入一个猫猫结构体
func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	//判断是不是添加第一支猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		//因为此时后面已经没有节点了，所以此时在这里选择将本节点的next指向自己
		//行成一个环形
		head.next = head
		fmt.Println(newCatNode.no, "号加入到环形列表中")
		return
	}
	//如果是第二只猫
	//先定义一个临时变量
	temp := head
	for {
		if temp.next == head {
			break
		}
		temp = temp.next
	}
	//加入到链表中
	temp.next = newCatNode
	newCatNode.next = head
}

// 删除环形链表的一个节点
func DelCircleNode(head *CatNode, id int) *CatNode {
	/*
		删除环形链表的思路如下：
		1、先让temp指向head；
		2、新建一个helper,先让其指向环形列表的最后；
		3、让temp和要删除id进行比较，如果相同，则通过helper删除【这里必须考虑如果删除的是头节点怎么办】；
	*/

	temp := head
	helper := head
	if temp.next == nil {
		//空链表的情况
		fmt.Println("这是一个空的环形链表，无法删除")
		return head
	}
	//如果只有一个节点
	if temp.next == head {
		temp.next = nil
		return head
	}
	//将helper定位到链表的最后
	for {
		if helper.next == nil {
			break
		}
		helper = helper.next
	}

	//有两个以上的节点
	flag := true
	for {
		if temp.next == head { //说明已经比较到最后一个了，还没有找到
			break
		}
		if temp.no == id { //说明找到了
			if temp == head {
				//说明删除的是头节点
				head = head.next
			}
			helper.next = temp.next
			fmt.Print(id, "号猫猫已经被删掉了\n")
			flag = false
			break
		}
		temp = temp.next     //移动【价值在于：不断得移动节点用该指针进行比较】
		helper = helper.next //移动【价值在于：一旦找到要删除的节点就使用该节点删除】
	}
	//这里还要比较一次
	if flag { //如果flag为真，说明上面代码一定没有节点删除成功
		if temp.no == id {
			helper.next = temp.next
			fmt.Print(id, "号猫猫已经被删掉了\n")
		} else {
			fmt.Println(id, "号猫猫删除失败")
		}
	}
	return head
}

// 输出环形列表
func listCircleList(head *CatNode) {
	temp := head
	if temp.next == nil {
		fmt.Println("CircleList is empty")
		return
	}

	for {
		fmt.Printf("[%d    %s]==>", temp.no, temp.name)
		//判断是否到头
		if temp.next == head {
			break
		}
		temp = temp.next
	}

}

func main() {
	fmt.Println("环形链表的使用！")

	//初始化一个环形链表的头节点
	head := &CatNode{}

	//创建一只新猫
	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}

	cat2 := &CatNode{
		no:   2,
		name: "jack",
	}
	cat3 := &CatNode{
		no:   3,
		name: "ross",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	fmt.Println("删除前：")
	listCircleList(head)
	head = DelCircleNode(head, 2)
	fmt.Println("删除后：")
	listCircleList(head)

}
