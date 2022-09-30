package main

import (
	"fmt"
	"os"
)

/*
Hash表的基本介绍
散列表是根据关键值而直接进行访问的数据结构。
也就是说，它通过把关键字映射到表中一个位置来访问记录，以加快查找速度。
这个映射函数叫做散列函数，存放记录的数组叫做散列表
*/

/*
实际案列：有一个公司，当有新的员工来报道时，要求将员工的信息加入（id，性别，年龄，地址），
当输入该员工的年龄的时候，要求查找该员工的所有信息
1.不使用数据库，尽量节省内存，速度越快越好  ==》 哈希表；
2.添加时，保证按照id从低到高插入；
*/

/*
思考分析：
1.使用链表来实现哈希表，该链表不带表头；【链表的第一结点就存放雇员信息】
2.用极少的内存却能保存大量的用户
*/

// 定义Emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

func (emp *Emp) ShowMe() {
	fmt.Printf("链表%d 找到该雇员 %d号", emp.Id%7, emp.Id)
}

// 定义EmpLink
// 这里的EmpLink不带表头，即第一个就存放雇员
type EmpLink struct {
	Head *Emp
}

// 方法待定:添加员工的方法
// 1.添加员工的方法，保证添加的时候，编号从小到大
func (el EmpLink) Insert(emp *Emp) {
	//创建辅助指针
	cur := el.Head
	var pre *Emp = nil
	if cur == nil {
		el.Head = emp
		return
	}
	//如果不是一个空链表，给temp找到相应的位置并插入
	//思路：让cur和emp比较，然后pre始终保持在cur前面

	for {
		if cur != nil {
			if cur.Id > emp.Id {
				//已经找到了位置
				break
			}
			pre = cur //保证同步
			cur = cur.Next
		} else {
			break
		}
	}

	//检查是否插入到链表的最后
	pre.Next = emp
	emp.Next = cur
}

// 定义hashtable，含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// 给hashTable编写Insert方法
func (ht *HashTable) Insert(emp *Emp) {
	//使用散列函数确定将该雇员确定加入到哪个链表
	linkNo := ht.HashFun(emp.Id)
	//使用对应的链表进行添加
	ht.LinkArr[linkNo].Insert(emp)

}

// 散列函数
func (ht *HashTable) HashFun(id int) int {
	return id % 7 //得到一个值，就是对应于链表的下标
}

// 编写EmpLink的查找方法
func (el EmpLink) FindById(no int) *Emp {
	cur := el.Head
	for {
		if cur != nil && cur.Id == no {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

// 显示所有员工,显示hashTable所有的员工
func (ht *HashTable) ShowAll() {
	for i := 0; i < len(ht.LinkArr); i++ {
		ht.LinkArr[i].ShowLink(i)
	}
}

// 显示当前链表的信息
func (el *EmpLink) ShowLink(no int) {
	if el.Head == nil {
		fmt.Printf("链表%d没有员工\n", no)
		return
	}

	//变量当前的链表，并显示数据
	cur := el.Head
	for {
		if cur != nil {
			fmt.Printf("链表%d ：雇员 id=%d name=%s==>", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

// 编写HashTable一个方法，完成查找具体是哪条链
func (ht *HashTable) FindById(id int) *Emp {
	//使用散列函数，确定将雇员应该放在哪个链表
	linkNo := ht.HashFun(id)
	return ht.LinkArr[linkNo].FindById(id)
}

func main() {
	id := 0
	name := ""
	key := ""
	//构建一个HasdTable
	var hashTable HashTable
	for {
		fmt.Println("==========================雇员系统=================")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show  表示展示所有雇员")
		fmt.Println("find  表示查找雇员")
		fmt.Println("exit  表示退出系统")
		fmt.Println("请输入你的选择：")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("请输入ID：")
			fmt.Scanln(&id)
			fmt.Println("请输入Name：")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashTable.Insert(emp)
		case "show":
			hashTable.ShowAll()
		case "find":
			fmt.Println("请输入id号：")
			fmt.Scanln(&id)
			emp := hashTable.FindById(id)
			if emp == nil {
				fmt.Println("该雇员不存在！")
			} else {
				fmt.Println("显示该雇员的信息！")
				emp.ShowMe()
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("未知输入，请重新添加")
		}
	}

}
