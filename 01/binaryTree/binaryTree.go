package main

import "fmt"

// 构建一个结构体
type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

/*
二叉树的前序 中序 后序遍历
*/

// 前序遍历：【先输出root节点，在输出左结点，再输出右节点】
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历：【先输出root节点的左子树，再输出root，最后输出右节点】
func InfixOrder(node *Hero) {
	if node != nil {
		PreOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Right)
	}
}

// 后序遍历：【先输出root节点的右子树，再输出root，最后输出左节点】
func PostOrder(node *Hero) {
	if node != nil {
		PreOrder(node.Left)
		PreOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)

	}
}

func main() {

	//构建一个二叉树
	root := &Hero{
		No:   1,
		Name: "宋江",
	}
	left1 := &Hero{
		No:   2,
		Name: "无用",
	}
	right1 := &Hero{
		No:   3,
		Name: "卢俊义",
	}
	//连接上面3个点
	root.Left = left1
	root.Right = right1
	right2 := &Hero{
		No:   4,
		Name: "林冲",
	}

	right1.Right = right2
	fmt.Println("前序遍历为：")
	PreOrder(root)
	fmt.Println("中序遍历为：")
	InfixOrder(root)
	fmt.Println("后序遍历为：")
	PostOrder(root)
}
