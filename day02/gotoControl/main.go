package main

import "fmt"

// goto

func main(){
	// 跳出多层for循环,推荐使用
	//flag := false
	//for i:=0;i<10;i++ {
	//	for j:='A';j<'Z';j++{
	//		if j=='C' {
	//			flag = true
	//			break // 跳出内层循环
	//		}
	//		fmt.Printf("%v-%c\n", i, j)
	//	}
	//	if flag {
	//		break // 跳出外层循环
	//	}
	//}
	// goto+label 跳出多层循环,尽量少用
	for i:=0;i<10;i++ {
		for j:='A';j<'Z';j++{
			if j=='C' {
				goto breakTag // 跳转到我指定的那个标签
			}
			fmt.Printf("%v-%c\n", i, j)
		}
	}
	breakTag:
		fmt.Println("over")
}
