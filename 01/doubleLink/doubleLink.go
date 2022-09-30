package main

import (
	"fmt"
)

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode //表示指向前一个英雄
	next     *HeroNode //表示指向下一个英雄
}

// 给双向链表添加一个节点
// 编写第一种方式：直接在单链表最后加入【简单】
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	//思路：
	//1、先找到该链表最后的那个节点
	//2、创建一个辅助节点：跑龙套，帮忙
	temp := head
	for {
		if temp.next == nil { //表示找到最后了
			//保证了退出for循环的时候temp一定是指向最后一个节点的【关键】
			break
		}
		temp = temp.next //不断移动temp指向下一个节点
	}
	//3、将newHeroNode加入到链表的最后
	temp.next = newHeroNode
	newHeroNode.pre = temp
}

// 给双向链表编写第二种插入方式：要考虑no的顺序来插入元素
func InsertHeroNode2(head *HeroNode, newHeroNode *HeroNode) {
	//思路：
	//1、先找到该链表中适当的节点
	//2、创建一个辅助节点：跑龙套，帮忙
	temp := head
	flag := true
	//让插入的节点no和temp的下一个节点的no做比较
	for {
		if temp.next == nil {
			//说明到链表的最后
			break
		} else if temp.next.no >= newHeroNode.no {
			//说newHeroNode应该插入temp后面
			break
		} else if temp.next.no == newHeroNode.no {
			//说明我们的链表中已经有这个no就不插入
			flag = false
			break
		}
		temp = temp.next
	}

	if !flag {
		fmt.Println("对不起，已经有这个no=", newHeroNode.no)
		return
	} else {
		newHeroNode.next = temp.next //ok
		newHeroNode.pre = temp       //ok
		if temp.next != nil {
			temp.next.pre = temp //ok
		}
		temp.next = newHeroNode //ok
	}
}

// 双向链表删除一个节点
func DeleteNode(head *HeroNode, id int) {
	//先创建一个临时节点
	temp := head
	//默认没有找到
	flag := false
	for {
		if temp.next == nil { //说明到链表的最后了
			break
		} else if temp.next.no == id {
			flag = true //找到了相等的元素
			break
		}
		temp = temp.next
	}
	if flag {
		temp.next = temp.next.next //ok
		if temp.next != nil {
			temp.next.pre = temp //ok,这里一定要小心
		}

	} else {
		fmt.Println("sorry,要删除的id不存在")
	}
}

// 显示链表的所有节点的信息
func listHeroNode(head *HeroNode) {
	//1、创建一个辅助节点
	temp := head
	//2、先判断一个链表是不是一个空链表
	if temp.next == nil {
		fmt.Println("list is empty")
		return
	}
	//3、遍历显示该链表
	for {
		fmt.Printf("[%d        %s     %s]==>", temp.next.no, temp.next.name,
			temp.next.nickname)
		//判断是否到链表头部了
		temp = temp.next
		if temp.next == nil {
			break
		}
	}
	fmt.Println()
}

// 显示链表的所有节点的信息
func listHeroNode2(head *HeroNode) {
	//1、创建一个辅助节点
	temp := head
	//2、先判断一个链表是不是一个空链表
	if temp.next == nil {
		fmt.Println("list is empty")
		return
	}
	//3、先让temp到链表的最后
	for {
		if temp.next == nil {
			break
		}
		temp = temp.next
	}

	//4、反向遍历显示该链表
	for {
		fmt.Printf("[%d        %s     %s]==>", temp.no, temp.name,
			temp.nickname)
		//判断是否到链表尾部
		temp = temp.pre
		if temp.pre == nil {
			break
		}
	}
	fmt.Println()
}

func main() {
	//1、先创建一个头节点
	head := &HeroNode{}

	//2、创意一个新的heroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
		//这里因为下一个英雄还没有确定，所以节点处先不用设置
	}

	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
		//这里因为下一个英雄还没有确定，所以节点处先不用设置
	}

	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
		//这里因为下一个英雄还没有确定，所以节点处先不用设置
	}
	hero4 := &HeroNode{
		no:       4,
		name:     "吴用",
		nickname: "智多星",
		//这里因为下一个英雄还没有确定，所以节点处先不用设置
	}

	InsertHeroNode(head, hero1)
	InsertHeroNode(head, hero2)
	InsertHeroNode(head, hero3)
	InsertHeroNode(head, hero4)
	fmt.Println("顺序打印如下所示：")
	listHeroNode(head)
	fmt.Println("逆序打印所下所示：")
	listHeroNode2(head)
}
