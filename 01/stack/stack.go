package main

import (
	"errors"
	"fmt"
)

/*
栈的应用场景
1、子程序的调用：在跳往子程序前，会将下一个指令的地址存放到栈中，直到子程序执行完毕后再将地址取出来，回到原来的程序中；
2、处理递归调用：和子程序的调用类似，只是除了存储的下一个指令的地址外，也将参数、区域变量等数据存入堆栈中；
3、表达式的转换与求值；
4、二叉树的遍历；
5、图形的深度优先搜索法（depth-first）；
*/

// 使用数组来模拟数组的使用
type Stack struct {
	MaxTop int //表示我们的栈最大的存放数的个数
	Top    int //表示栈顶，因为栈顶是固定的，因此我们直接使用Top
	arr    [5]int
}

// 入栈
func (s *Stack) Push(val int) error {
	//先判断栈是否已经满了
	if s.Top == s.MaxTop-1 {
		fmt.Println("stack full")
		return errors.New("stack full")
	}
	s.Top++
	s.arr[s.Top] = val
	return nil
}

// 出栈
func (s *Stack) Pop() (val int, err error) {
	//先判断栈是否已经空了
	if s.Top == -1 {
		fmt.Println("stack empty")
		return 0, errors.New("stack empty")
	}
	//先取值
	val = s.arr[s.Top]
	s.Top--
	return val, nil
}

// 遍历栈：从栈顶开始遍历
func (s *Stack) List() {
	//先判断栈是否已经满了
	if s.Top == -1 {
		fmt.Println("栈空了")
	}
	fmt.Println("栈的情况如下所示：")
	for i := s.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, s.arr[i])
	}
}

func main() {

	s := &Stack{
		MaxTop: 5,  //表示最多存放5个数到栈中
		Top:    -1, //当Top等于-1表示空
	}
	s.Push(1)
	s.Push(5)
	s.Push(10)
	s.Push(12)
	s.Push(23)
	s.List()
	fmt.Println("出栈情况如下")
	s.Pop()
	s.List()

}
