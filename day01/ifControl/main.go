package main

import "fmt"

// if 判断
func main(){
	age := 19
	if age > 18 {
		fmt.Println("年龄大于18岁")
	} else {
		fmt.Println("小于18岁")
	}
	// 多个判断
	if age > 35 {
		fmt.Println("中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("未成年")
	}
	if age1 := 19; age1>18 {	// age1变量只在if条件判断语句中生效
		fmt.Println("成年")
	} else {
		fmt.Println("未成年")
	}
	//fmt.Println(age1)  // age1声明再if作用域内,在if作用域外不能找到
}
