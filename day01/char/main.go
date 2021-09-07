package main

import (
	"fmt"
	"unicode"
)

func main() {
	s := "hello 这个世界"
	// 中文或其他叫rune类型
	// 英文或字符叫byte类型, uint8 类型
	n := len(s)  // 求字符串的长度,将长度保存到变量n中
	fmt.Println(n)

	//for i := 0; i < len(s); i++ {
	//	fmt.Println(s[i])
	//	//fmt.Printf("%c\n", s[i])  // %c 字符
	//}
	for _, c := range s {
		fmt.Println(c)  // 从字符串中拿出具体的字符
		fmt.Printf("%c\n", c)
	}
	// 字符串修改
	s2 := "户罗布"
	s3 := []rune(s2) // 把字符串强制转换成一个rune切片
	s3[0] = '红'
	fmt.Println(string(s3))  // 把rune切片强制转为字符串
	c1 := "红"	// string类型
	c2 := '红'   // rune类型
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "H"	// string类型
	c4 := 'H'   // rune类型  c4 := byte('H')
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	// 类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Printf("f:%T  f=%v\n", f, f)
	count := 0
	for _, c := range s{
		if unicode.Is(unicode.Han, c){
			count++
		}
	}
	fmt.Printf("\"%s\"中有中文字符[%v]个.\n", s, count)
}
