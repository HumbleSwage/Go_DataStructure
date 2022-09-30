package main

import (
	"errors"
	"fmt"
	"strconv"
)

/*
如何使用栈来实现一个对表达式的计算
1、创建两个栈numStack和operStack，分别用来存放运算元素和运算符号；
2、index := 0；
3、计算表达式，计算一个字符串；
4、如果扫描发现是一个数字，则直接入numStack；
5、如果发现是一个运算符；

	（1）如果运算符栈是一个空栈，则该符号直接入栈；
	（2）如果发现operStack当前栈顶符号的运算符的优先级大于等于当前准备入栈的运算符的优先级，则需要计算对应位置的运算元素被从新压入栈；

6、如果扫描表达式完毕，依次从符号位取出符号，然后从数栈取出两个数进行运算；
*/

// 使用数组来模拟数组的使用
type Stack struct {
	MaxTop int //表示我们的栈最大的存放数的个数
	Top    int //表示栈顶，因为栈顶是固定的，因此我们直接使用Top
	arr    [20]int
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

// 判断一个字符是不是运算符号
func (s *Stack) IsOper(val int) bool {
	if val == 42 || val == 43 || val == 45 || val == 47 {
		return true
	} else {
		return false
	}
}

// 运算的方法
func (s *Stack) Cal(num1 int, num2 int, oper int) int {
	res := 0
	switch oper {
	case 42:
		res = num1 * num2
	case 43:
		res = num1 + num2
	case 45:
		res = num2 - num1
	case 47:
		res = num2 / num1
	default:
		fmt.Println("未知运算符号")
	}
	return res
}

// 编写一个方法，返回某个运算符号的优先级别【程序定】
func (s *Stack) Priority(oper int) int {
	res := 0
	if oper == 42 || oper == 47 { //乘法或者除法
		res = 1
	} else if oper == 43 || oper == 45 { //加法或者减法
		res = 0
	}
	return res

}

func main() {
	//数字栈
	numStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	//符号栈
	operStack := &Stack{
		MaxTop: 20,
		Top:    -1,
	}
	exp := "3+2*6-2"

	//定义一个index帮助我们扫描
	index := 0

	//为了配合我们运算，我们定义需要的变量
	num1 := 0
	num2 := 0
	oper := 0
	result := 0
	keepNum := "" //初始化后面保存字符串的数组
	for {
		//针对多位数的计算，这里需要额外的逻辑
		ch := exp[index : index+1] //返回单个的字符串
		//传入的必须是字符对应的asii码，所以这里需要一定的转换技巧,先转换为byte切片，然后再取切片的第一个,再将byte转换为int
		temp := int([]byte(ch)[0])
		if operStack.IsOper(temp) { //说明是计算符
			if operStack.Top == -1 {
				operStack.Push(temp)
			} else {
				//这里有一个运算级别优先级别的比较
				if operStack.Priority(operStack.arr[operStack.Top]) >= operStack.Priority(temp) {
					//这个时候数栈里一定有元素，所以这个时候从数栈开始取元素
					num1, _ = numStack.Pop()
					num2, _ = numStack.Pop()
					oper, _ = operStack.Pop()
					result = operStack.Cal(num1, num2, oper)
					//再把结果压入numStack栈
					numStack.Push(result)
					//当前的符号压入operStack栈
					operStack.Push(temp)

				} else {
					operStack.Push(temp)
				}
			}
		} else {
			//说明是数
			//处理多位数的一个思路
			/*
				1.定义一个变量keepNum string,做字符串的拼接；
				2.每次要向index的前面的字符测试一下，看看是不是运算符，然后处理
			*/
			keepNum += ch
			//如果已经到最后，直接将keepNum
			if index == len(exp)-1 {
				val, _ := strconv.ParseInt(keepNum, 10, 64)
				numStack.Push(int(val))
			} else {
				//向index后面的后面一位探测，判断是不是运算符号
				if operStack.IsOper(int([]byte(exp[index+1 : index+2])[0])) {
					val, _ := strconv.ParseInt(keepNum, 10, 64)
					numStack.Push(int(val))
					keepNum = "" // 清空
				}
			}
			// val, _ := strconv.ParseInt(ch, 10, 64)
			// numStack.Push(int(val))
		}
		//继续扫描,先判断index是否已经到计算表达式的最后
		if index+1 == len(exp) {
			break
		}
		index++

	}
	//如果扫描表达式完毕，依次从符号栈取出符号，然后从数栈取出两个数
	//运算后的结果，入数栈，直到符号栈为空
	for {
		if operStack.Top == -1 {
			//退出条件
			break
		}
		num1, _ = numStack.Pop()
		num2, _ = numStack.Pop()
		oper, _ = operStack.Pop()
		result = operStack.Cal(num1, num2, oper)
		//再把结果压入numStack栈
		numStack.Push(result)
	}
	//如果我们的算法没有问题，那么numStack最终只会有我们的结果
	res, _ := numStack.Pop()
	fmt.Printf("%s 的结果为%d\n", exp, res)

}
