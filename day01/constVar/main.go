package main

import "fmt"

// 常量
// 定义常量之后不能修改
// 在程序运行中不会改变的量
const pi = 3.141592675

// 批量声明常量
const (
	statusOK = 200
	notFound = 404
)

// 批量声明常量时,如果某一行声明后没有赋值,默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)
// iota
const (
	a1 = iota // 0
	a2		  // 1
	a3
)
const (
	b1 = iota
	b2
	_
	b3
)
const (
	c = iota  //0
	c2 = 100  // 100
	c3 = iota  // 2
	c4		   // 3
)
// 多个变量声明在一行
const (
	d1, d2 = iota + 1, iota + 2   // d1=1,d2=2
	d3, d4 = iota + 1, iota + 2   // d3=2,d4=3
)
// 使用iota定义数量级
const(
	_ = iota
	KB = 1 << (10*iota)
	MB = 1 << (10*iota)
	GB = 1 << (10*iota)
	TB = 1 << (10*iota)
	PB = 1 << (10*iota)
)
func main(){

	//pi = 1334   pi常量不能再赋值
	fmt.Println(n1)
	fmt.Println(n2)
	fmt.Println(n3)
	// 打印iota
	//fmt.Println(a1)
	//fmt.Println(a2)
	//fmt.Println(a3)

	fmt.Println(b1)
	fmt.Println(b2)
	fmt.Println(b3)
}

