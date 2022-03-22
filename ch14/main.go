package main

import (
	"fmt"
)

func main(){
	var s string
	var sp *string
	fmt.Println("s:",s,"sp:",sp)
	fmt.Println("s:",&s,"sp:",sp)
	s = "cold"
	// *sp = "jun" 因为未分配内存所以赋值报错
	fmt.Println("s:",s,"sp:",sp)
	fmt.Println("\nnew函数")
	sp = new(string)
	fmt.Println("sp:",sp)
	*sp = "jun"
	fmt.Println("sp:",*sp)
	
	fmt.Println("\n变量初始化")
	p := person{name:"cold",age:18}
	fmt.Println("name:",p.name,"age:",p.age)
	
	fmt.Println("\n指针变量初始化")
	//pp := NewPerson()
	//fmt.Println("name:",pp.name,"age:",pp.age)
	ppp := NewPerson("cute",18)
	fmt.Println("name:",ppp.name,"age:",ppp.age)
}

type person struct{
	name string
	age int
}

/* func NewPerson() *person{
	p:=new(person)
	p.name="june"
	p.age = 18
	return p
}*/

func NewPerson(name string, age int) *person{
	p:=new(person)
	p.name=name
	p.age = age
	return p
}