package main

import (
	"fmt"
	"strings"
)

// 字符串
func main() {
	// \ 转义特殊字符
	path := "\"D:\\Go\\src\\code\\github\\studygo\""
	fmt.Printf(path)

	s := "I'm OK"
	fmt.Println(s)
	// 多行的字符串
	s2 := `
	人生苦短
		我用python
	let's go
	`
	fmt.Println(s2)
	s4 := `D:\Go\src\code\github\studygo`
	fmt.Println(s4)
	// 字符串相关操作
	fmt.Println(len(s4))

	// 字符串拼接
	name := "理想"
	world := "dsb"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Printf("%s%s", name, world)
	fmt.Println(ss1)
	// 分割
	ret := strings.Split(s4, "\\")
	fmt.Println(ret)
	// 包含
	fmt.Println(strings.Contains(ss, "理性"))
	//前缀
	fmt.Println(strings.HasPrefix(ss, "理想"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "理想"))

	s5 := "abcdeb0"
	fmt.Println(strings.IndexAny(s5, "c"))
	fmt.Println(strings.LastIndex(s5, "b"))

	// 拼接
	fmt.Println(strings.Join(ret, "//"))
}
