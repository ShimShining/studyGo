package main

import "fmt"

// for循环
func main() {
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	// 变换1
	//var i = 5
	//for ; i < 10; i++ {
	//	fmt.Println(i)
	//}
	// 变换2
	var i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}
	// 无线循环
	//for {
	//	fmt.Println("1234")
	//}
	// for range
	s := "hello 世界"
	for i,v := range s {
		fmt.Printf("index=%d, value=%c\n",i, v)
	}
	// 打印99 乘法表
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			mul := i*j
			if mul > 9 {
				fmt.Printf("%d * %d = %v ", j, i, mul)
			} else {
				fmt.Printf("%d * %d = %v  ", j, i, mul)
			}
		}
		fmt.Println()
	}
	// 倒序打印99乘法表
	for i := 9; i > 0; i-- {
		for j := 1; j <= i; j++ {
			mul := i*j
			if mul > 9 {
				fmt.Printf("%d * %d = %v ", j, i, mul)
			} else {
				fmt.Printf("%d * %d = %v  ", j, i, mul)
			}
		}
		fmt.Println()
	}
}
