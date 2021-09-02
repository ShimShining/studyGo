package main

import "fmt"

// Go语言变量命名推荐使用驼峰命名, 推荐使用小驼峰
// 声明变量
//var name string
//var age int
//var isOK bool

// 批量声明
var (
	firstName string
	beforeAge int
	isOk bool
)
func main() {
	firstName = "理想"
	beforeAge = 16
	isOk = true
	// Go语言中非全局变量声明后必须使用,不使用就编译不过去
	fmt.Print(isOk)  // 在终端打印内容,无换行
	fmt.Println()  // 打印空行
	fmt.Printf("name:%s", firstName) // %s是格式化占位符
	fmt.Println(beforeAge)  // 打印完成后自后会加一个换行符
	//var localVar string

	// 声明变量同时赋值
	var s1 string = "shimmer"
	fmt.Println(s1)

	// 类型推导
	var s2 = "shing"
	fmt.Printf("类型推导:%s", s2)

	// 简短变量声明, 只能在函数里用
	shortVar := "shortVar"
	fmt.Println(shortVar)

	//shortVar := "repeatDecalare"  // 同一个作用域{}不能重复声明同名的变量

	// 匿名变量 : _,匿名变量不占用命名空间,不会分配内存
}
