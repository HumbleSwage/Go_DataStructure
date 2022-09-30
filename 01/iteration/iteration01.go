package main

import (
	"fmt"
)

/*
1、迭代函数一般都是独立开栈的
2、以后是非常经典的一个迭代案列，请试着说出他们的输出结果的差异
*/
func test01(val int) {
	if val > 2 {
		val--
		test01(val)
	}
	fmt.Println(val)
}

func test02(val int) {
	if val > 2 {
		val--
		test02(val)
	} else {
		fmt.Println(val)
	}
}

func main() {
	fmt.Println("迭代的经典案例：")
	fmt.Println("test01的输出结果如下所示：")
	test01(4)
	fmt.Println("test02的输出结果如下所示：")
	test02(4)

}
