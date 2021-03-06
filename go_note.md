# go_note

参考博文 : [李文周的博客](https://www.liwenzhou.com/posts/Go/go_menu/)

## go语言开发环境

### 安装go开发包

#### 下载

### 配置GOPATH

### GO语言项目结构

### 下载安装VS code

#### 下载

#### 安装

#### 安装中文插件

#### 安装Go扩展

#### 常用介绍

### 编译

使用 go build

 	1. 在项目目录下执行
 	2. 在其他路径下执行,需要在后面加上项目的路径
 	3. go build -0 temp.exe 指定编译之后的文件名称

#### go run

像执行脚本文件一样执行go代码

#### go install

go install 分为两步 

 	1. 先编译得到一个可执行文件
 	2. 将可执行文件拷贝到GOPATH/bin

#### 跨平台编译

```
// Windows平台 cmd下编译一个能在linux下运行的程序
SET CGO_ENABLED=0 // 禁用CGO
SET GOOS=Linux // 目标平台是Linux
SET GOARCH=amd64 // 目标处理器架构师amd64
go build // 然后再执行go build命令得到的就是Linux平台运行的可执行文件
// mac
CGO_ENABLED=0 OGOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 OGOOS=windows GOARCH=amd64 go build
// linux
```



## Go语言基础

### Go 语言文件的基本结构

```go
package main

// 导入语句
import "fmt"

// 函数外只能放置标识符(变量|常量|函数|类型)的声明
// fmt.Println("Hello,GO") 放置go语句是非法
func main(){
	fmt.Println("Hello,GO")
}
```

###  变量和常量

#### 标识符与关键字

##### 标识符

程序员定义的具有特殊含义的单词,变量名,函数名,常量名,字母数字下划线组成,只能由字母和_开头

##### 关键字

编程语言中预先定义好的具有特殊含义的标识符,关键字和保留字都不建议作为变量名

```go
break default func interface...
```

#### 变量

##### 变量类型

go语言的变量必须先声明,再使用

##### 变量声明

```go
var s1 string // 声明一个保存字符串类型数据的变量s1
var age int 
```

##### 注意事项

1. 函数外的每个语句都必须以关键字开始(var,const,func等)
2. :=不能使用在函数外
3. _多用于占位,表示忽略值
4. 变量名不能重复声明

#### 常量

```go
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
```

#### 字符串

Go语言字符串是用双引号包裹的!!!!

Go语言中单引号包裹的是字符

```go
s := "hello go"  // 字符串
// 单独的字母,汉字,符号表示一个字符
c1 := 'C'
c2 := '1'
c3 := "字"
```

### 回顾day01

#### Go命令

`go build` : 编译go程序

`go build -o "xx.exe"`: 编译成xx.exe文件

`go run main.go`: 像脚本一样运行main.go文件

`go install`:先编译后拷贝

#### go语言基础语法

存放Go源代码的文件名后缀是.go

文件第一行:`package` 关键字声明包名

```go
// 单行注释

/*

多行注释

*/
```

编译一个可执行文件,必须要声明一个main包,必须要有main函数,没有参数没有返回值,main函数是入口

函数外的语句必须以关键字开头`var`,`const`

函数内部定义的变量必须使用

#### 变量

声明方式有三种

1. `var name string`
2. `var name = "hello 好"`
3. 函数内部专属: `name := "hello 王子"`

匿名变量(哑元变量): 有些数据必须用变量接受,但是又不想使用,使用_接受

#### 常量

`const PI = 3.1415926`

`const UserNotExistErr = 10000`

iota: 实现枚举

2个要点:

1. iota在const关键字出现时被重置为0
2. 每新增一行常量声明,iota累加1

#### 流程控制

##### if

```go
var age = 19
if age > 18 {
    fmt.Println(age)
} else if age > 7 {
     fmt.Println(age大于7)
} else {
     fmt.Println(age小于7)
}
```

##### for

1.  标准for循环

   ```go
   for i := 0; i < 10; i++ {
       fmt.Println(i)
   }
   ```

2. 循环变量声明在外面

   ```go
   var i = 0
   for ; i < 10; i++ {
       fmt.Println(i)
   }
   ```

3. for结束语句在里面

   ```go
   var i = 0
   for i < 10{
       fmt.Println(i)
       i++ 
   }
   ```

4. 无限循环

   ```go
   for {
       fmt.Println("无限循环")
   }
   ```

5. for range

   ```go
   var s = "hello 这个世界"
   for _, v := range s {
       fmt.Println(v) // 打印出来是ASCII码
       fmt.Printf("%d %c\n", i, v)
   }
   ```

#### 基本数据类型

整型

​	无符号整型: `uint8`,`uint16`,`uint32`, `uint64`

​	带符号整型:`int8`,`int16`,`int32`, `int64`

​	`int`和`uint`: 具体是32位还是64位看操作系统

​	`uintptr`: 表示指针

其他进制

​	go语言中无法直接定义二进制

​	可以定义八进制和16进制

```go
var n1 = 0777 // 8进制
var n2 = 0xff // 16 进制
fmt.Printf("%o\n", n1)  // 打印八进制
fmt.Printf("%x\n", n2)	// 打印十六进制
```



浮点型

​	`float32`

​	`float64`

​	Go语言默认的浮点数是`float64`

复数

`complex128`和`complex64`

布尔值

​	true和false

​	不能和其他的类型做转换

字符串

​	常用方法

​	字符串不能修改

byte和rune类型

​	都属于类型别名

字符串,字符,字节

​	字符串: 双引号包裹的是字符串

​	字符: 单引号包含的字符,单个字母,单个符号,单个文字

​	字节: 1byte=8 bit

​	go语言中字符串都是UTF-8编码,utf8编码中一个常用汉字一般占3个字节
