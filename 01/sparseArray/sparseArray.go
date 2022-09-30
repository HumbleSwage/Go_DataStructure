package main

import (
	"fmt"
	"encoding/json"
	"os"
	"io/ioutil"
)


type ValNode struct {
	Row int
	Column int
	Value int //如果存在其他类型这里也能够使用interface{}类型
} 

//实现对稀疏矩阵的保存
func main(){
	//1、构造稀疏矩阵
	var sparseArray [11][11]int
	sparseArray[1][2] = 1
	sparseArray[2][3] = 2

	//2、打印稀疏矩阵
	fmt.Println("构造的稀疏矩阵为：")
	for _,i  := range sparseArray {
		for _,j := range i {
			fmt.Printf("%d\t",j)
		}
		fmt.Println()
	}

	//3、通过特定的形式来保存稀疏矩阵，算力换空间
	var NodeList []ValNode
	var  validElementsNum int //用于记录元素的有效个数
	for _, i := range sparseArray {
		for _, j := range i {
			if j != 0 {
				validElementsNum++
			}
		}
	}
	if validElementsNum > len(sparseArray) * len(sparseArray[0])/5{
		fmt.Println("该稀疏矩阵不够稀疏，使用本方法没有办法提高效率！")
		return
	}
	NodeList = append(NodeList,ValNode{
		Row:len(sparseArray),
		Column:len(sparseArray[0]),
		Value:validElementsNum,
	})

	for index1 , i := range sparseArray {
		for index2 , j := range i {
			if j != 0 {
				
				NodeList = append(NodeList,ValNode{
					Row:index1,
					Column:index2,
					Value:j,
				})
			}
		}
	}
	//4、将其保存在文件中
	fmt.Println("以指定的形式来保存稀疏矩阵：")
	for _,valNode := range NodeList {
		fmt.Printf("%d\t%d\t%d\n",valNode.Row,valNode.Column,valNode.Value)
	}
	data,err := json.Marshal(NodeList)
	if err != nil {
		fmt.Println("json.Marshal err=",err)
		return
	}
	dstFile,err := os.Create("info.dat")
 	if err!=nil {
        fmt.Println("os.Create err=",err)  
        return
    } 
    defer dstFile.Close()
    dstFile.WriteString(string(data) + "\n")
    //5、从文件中读取，将将其转换为原始稀疏矩阵
    Rfile, err := os.Open("info.dat")
    if err != nil {
        fmt.Println("read file fail", err)
        return
    }
    defer Rfile.Close()
    fd, err := ioutil.ReadAll(Rfile)
    if err != nil {
        fmt.Println("read to fd fail", err)
        return
    }
    var NodeListAfter []ValNode
    json.Unmarshal(fd,&NodeListAfter)
	 //    row := NodeListAfter[0].Row
		// column := NodeListAfter[0].Column
	var NewSparseArray [11][11]int
    for index,valNode := range NodeListAfter {
    	if index == 0 {
    		continue
    	}
    	NewSparseArray[valNode.Row][valNode.Column] = valNode.Value
    }
	fmt.Println("从新的数据结构到稀疏矩阵为：")
	for _,i  := range NewSparseArray {
		for _,j := range i {
			fmt.Printf("%d\t",j)
		}
		fmt.Println()
	}










}