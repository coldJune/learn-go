package main

import "fmt"

func main(){
	fmt.Println("if条件语句")
	
	i:= 5
	if i>10{
		fmt.Println("i>10")
	}else if i>5&&i<=10{
		fmt.Println("i>5&&i<=10")
	}else{
		fmt.Println("i<=5")
	}
	
	if s:="hello"; i>10{
		fmt.Println(s,"i>10")
	}else if i>5&&i<=10{
		fmt.Println(s,"i>5&&i<=10")
	}else{
		fmt.Println(s,"i<=5")
	}
	
	fmt.Println("\nswitch")
	switch i:=1;{
		case i>10:
			fmt.Println("i>10")
		case i>5 && i<=10:
			fmt.Println("i>5 && i<=10")
		default:
			fmt.Println("i<5")
	}
	
	fmt.Println("\nswitch fallthrough")
	switch i:=1;{
		case i<10:
			fmt.Println("i<10")
			fmt.Println(i>5 && i<=10)
			fallthrough // 不会判断下一个case的条件 而是直接执行
		case i>5 && i<=10:
			fmt.Println("i>5 && i<=10")
		default:
			fmt.Println("i<5")
	}
	
	fmt.Println("\nfor循环")
	
	sum:=0
	for i:=1;i<=100;i++{
		sum +=i
	}
	fmt.Println("sum=",sum)
	
	i=1
	for {
		sum +=i
		i++
		if i>100{
			break
		}
	}
	fmt.Println("sum=",sum)

}