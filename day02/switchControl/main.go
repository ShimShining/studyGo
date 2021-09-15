package main

import "fmt"

// switch
// 简化大量的判断
func main() {
	//var n = 5
	//if n == 1{
	//	fmt.Println("大")
	//} else if n==2{
	//	fmt.Println("是指")
	//} else if n==3 {
	//	fmt.Println("3")
	//} else if n == 4 {
	//	fmt.Println("4")
	//} else if n == 5 {
	//	fmt.Println("5")
	//} else {
	//	fmt.Println("other")
	//}
	// switch 简化上面的代码
	switch n:=3; n {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("是指")
	case 3:
		fmt.Println("终止")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小智")
	default:
		fmt.Println("other")
	}

	// 匹配多个中一个
	switch n:=3; n {
	case 1,3,5,7,9:
		fmt.Println("奇数")
	case 2,4,6,8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
	// case 后加判断
	age := 30
	switch {
	case age < 22 :
		fmt.Println("好好学习")
	case age >= 22 && age < 60:
		fmt.Println("好好工作")
	case age >= 60:
		fmt.Println("好好享受")
	default:
		fmt.Println("how")
	}
	// fallthrough 可以执行满足条件case的下一个case,是为了兼容C语言中的case设计的 不建议使用
}
