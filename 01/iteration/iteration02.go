package main

import "fmt"

/*
使用迭代解决迷宫问题
*/

// 编写一个函数，完成老鼠找迷宫
// 地图必须保证始终是同一个*[8][7]int，所以必须是引用
// i,j表示对地图上哪个点进行测试
func FindWay(myMap *[8][7]int, i int, j int) bool {
	//分析什么情况应该下，应该找到出路
	//myMap[6][5] = 2
	if myMap[6][5] == 2 {
		return true
	} else {
		//说明需要继续寻找
		if myMap[i][j] == 0 { // 如果这个点可探测
			//假设这个点是可以通的，但是需要探测，上下左右
			myMap[i][j] = 2
			if FindWay(myMap, i+1, j) { //下
				return true
			} else if FindWay(myMap, i, j+1) { //右
				return true
			} else if FindWay(myMap, i-1, j) { //上
				return true
			} else if FindWay(myMap, i, j-1) { //左
				return true
			} else { // 上下走不通
				myMap[i][j] = 3
				return false
			}

		} else { //说明这个点不能探测，为1，是墙
			return false
		}
	}
}

func main() {
	//先创建一个二维数组，模拟迷宫
	//规则如下所示
	//1.如果元素值为0，就是墙；
	//2.如果元素值是1，就是没有走过的墙；
	//3.如果元素的值2，是另一个通路；
	//4.如果元素的值3，就是走过的墙；
	var myMap [8][7]int
	//先把地图的最上和最下设置为1
	for i := 0; i < 7; i++ {
		myMap[0][i] = 1
		myMap[7][i] = 1
	}
	myMap[3][1] = 1
	myMap[3][2] = 1
	//地图的最左边和最右边设置为1
	for i := 0; i < 7; i++ {
		myMap[i][0] = 1
		myMap[i][6] = 1
	}

	//使用测试
	FindWay(&myMap, 1, 1)
	fmt.Println("探测完毕！")
	//输出地图
	for i := 0; i < 8; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(myMap[i][j], " ")
		}
		fmt.Println()
	}

}
