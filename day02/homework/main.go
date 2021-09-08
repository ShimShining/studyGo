package main

import "fmt"

//byte和rune

func main() {
	s1 := "hello"
	s2 := "上海自来水来自海上"

	for _,v := range s1 {
		fmt.Printf("%c   %T\n", v, v)
	}

	for _, v := range s2 {
		fmt.Printf("%c    %T\n", v, v)
	}
}
