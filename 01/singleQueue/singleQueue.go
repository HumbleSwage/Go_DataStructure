package main

import (
	"fmt"
	"os"
	"errors"
)

type Queue struct {
	maxSize int
	array [5]int
	head int
	tail int
}
//注意这个情况下首个位置是没有元素的希望+
//添加数据到队列当中
func (queue *Queue)AddQueue(value int)(err error){
	//先判断队列是否已满
	if queue.tail == queue.maxSize - 1 {//重要的提示；tail是队列的尾部（含最后的元素）
		return errors.New("queue full")
	}
	//注意以下两句话的逻辑顺序
	queue.tail++
	queue.array[queue.tail] = value
	return
}


//显示队列，找到队首
func (queue *Queue)ShowQueue(){
	fmt.Println("队列当前的情况是：")
	//queue.head不包含队首元素
	for i := queue.head + 1 ; i <= queue.tail ; i++ {
		fmt.Printf("array[%d]=%d\t",i,queue.array[i])
	}
	fmt.Println()
}


//从队列中取出数据
func (queue *Queue)GetQueue()(val int,err error){
	//首先判断队列是否为空
	if queue.tail == queue.head {
		return -1,errors.New("queue empty")
	}
	queue.head++
	val =  queue.array[queue.head]
	return val,nil
}





//实现链表
func main(){
	//先创建一个队列
	queue := Queue {
		maxSize : 5,
	}
	for {
		fmt.Println("1、输入add表示添加数据到队列;")
		fmt.Println("2、输入get表示展示数据队列;")
		fmt.Println("3、输入show表示显示数据队列;")
		fmt.Println("4、输入exit表示退出;")
		fmt.Println(" 请输入你的选择：")
		var key string
		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入你要入队的数据：")
			var val int
			fmt.Scanf("%d",&val)
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println(err)
			}
		case "get":
			_,err := queue.GetQueue()
			if err != nil {
				fmt.Println(err)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
			
		default:
			/* code */
			return
		}

		
	}


}