package main

import "fmt"
import "strconv"
import "Strings"

func main(){
	// 整型
	fmt.Println("整型")
	var i int = 9
	fmt.Println(i)
	
	var j = 10
	fmt.Println(j)
	
	var(
		n int = 0
		k int = 2
	)
	fmt.Println(n)
	fmt.Println(k)
	
	// 浮点数
	fmt.Println("浮点数")
	var f32 float32 = 1.1
	var f64 float64 = 3.141592
	fmt.Println("f32=",f32,"f64=",f64)
	
	// 布尔型
	fmt.Println("布尔型")
	var bf bool = false
	var bt bool = true
	fmt.Println("bf is",bf,"bt is",bt)
	
	// 字符串
	fmt.Println("字符串")
	var s1 = "hello"
	var s2 = "world"
	fmt.Println("s1 is",s1,"s2 is",s2)
	fmt.Println("s1+s2=",s1+s2)
	
	// 默认值
	fmt.Println("默认值")
	var zi int
	var zf float64
	var zb bool
	var zs string
	fmt.Println(zi,zf,zs,zb)
	
	// 变量简写
	fmt.Println("变量简写")
	is:=10
	bfs:=false
	s1s:="hello"
	fmt.Println(is,bfs,s1s)
	
	// 指针
	fmt.Println("指针")
	pi:=&is
	fmt.Println(*pi)
	fmt.Println(pi)
	
	// 赋值
	fmt.Println("赋值")
	fmt.Println("before",i)
	i = 20
	fmt.Println("after",i)
	
	// 常量
	fmt.Println("常量")
	const name = "coldjune"
	fmt.Println(name)
	
	// iota 常量生成器
	fmt.Println("常量生成器")
	const(
		one = iota + 1
		two
		three
		four
	)
	fmt.Println(one,two,three,four)
	
	// 字符串和数字互转
	fmt.Println("字符串和数字互转")
	i2s:=strconv.Itoa(i)
	s2i,err:=strconv.Atoi(i2s)
	fmt.Println(i2s,s2i,err)
	
	i2f:=float64(i)
	f2i:=int(f64)
	fmt.Println(i2f,f2i)

	// Strings包
	fmt.Println("Strings包")
	fmt.Println("s1:",s1,"前缀是否是H",strings.HasPrefix(s1,"H"))
	fmt.Println("s1:",s1,"查找字符串o",strings.Index(s1,"o"))
	fmt.Println("s1:",s1,"全部转为大写",strings.ToUpper(s1))
}