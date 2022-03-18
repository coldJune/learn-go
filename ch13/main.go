package main

import (
	"fmt"
)

type address struct{
	province string
	city string
}
func (addr address)String() string{
	return fmt.Sprintf("the addr is %s%s",addr.province, addr.city)
}
func main(){
	add := address{province:"SC",city:"CD"}
	printString(add)
	printString(&add)
	
	var si fmt.Stringer = address{province:"SC",city:"CD"}
	printString(si)
	// sip := &si 虽然指向具体类型的指针可以实现一个接口，但是指向接口的指针永远不可能实现该接口
	//printString(sip)
	
	fmt.Println("\n修改参数")
	p :=person{name:"cold",age:26}
	fmt.Printf("main函数：p的内存地址为%p\n",&p)
	modifyPerson(p)
	fmt.Println("person name:",p.name,"age:",p.age)
	modifyPersonP(&p)
	fmt.Println("person name:",p.name,"age:",p.age)
	fmt.Println("\n引用类型")
	m := make(map[string]int)
	m["cold"] = 18
	fmt.Println("cold的年龄为:",m["cold"])
	fmt.Printf("main函数：m的内存地址为%p\n",m)
	modifyMap(m)
	fmt.Println("cold的年龄为:",m["cold"])
	
}

func printString(s fmt.Stringer){
	fmt.Println(s.String())
}

type person struct{
	name string
	age int
}

func modifyPerson(p person){
	fmt.Printf("modifyPerson函数：p的内存地址为%p\n",&p)
	p.name = "june"
	p.age = 18
}

func modifyPersonP(p *person){
	fmt.Printf("modifyPersonP函数：p的内存地址为%p\n",p)
	p.name = "june"
	p.age = 18
}

func modifyMap(p map[string]int){
	fmt.Printf("modifyMap函数：m的内存地址为%p\n",p)
	p["cold"] = 26
}