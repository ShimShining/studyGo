package main

import "fmt"

// 浮点数 布尔值

func main() {
	//math.MaxFloat32 // float32 最大值
	f1 := 1.2345
	fmt.Printf("%T\n", f1) // go语言中小数默认是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) // 显示声明float32类型
	// f1 = f2	// float32类型的值不能直接赋值给float64类型

	// 布尔
	b1 := true
	var b2 bool   // 布尔值默认值为false
	fmt.Printf("%T  value:%v\n", b1, b1)
	fmt.Printf("%T  value:%v\n", b2, b2)
}
