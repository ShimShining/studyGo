package main

import "fmt"

// 流程控制之跳出for循环

func main() {
	// 当i=5 跳出for循环
	for i:=0;i<10;i++ {
		if i == 5{
			break
		}
		fmt.Println(i)
	}
	fmt.Println("over for")
	// 当i=5 跳过当次for循环
	for i := 0;i<10;i++{
		if i==5 {
			fmt.Println("skip")
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("over")
}
