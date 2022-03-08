package main

import "fmt"


type person struct{
	name string
	age uint
}

func main(){
	fmt.Println("结构体")
	var p person
	fmt.Println(p.name,p.age)
	
	p = person{"jun",18}
	fmt.Println(p.name,p.age)
	
	p = person{name:"june",age:26}
	fmt.Println(p.name,p.age)
	
	p = person{age:16}
	fmt.Println(p.name,p.age)
	
	pe := people{
		name:"june",
		addr:address{
			city:"cd",
		},
		address:address{
			city:"bj",
		},
	}
	fmt.Println(pe.addr.city)
	
	fmt.Println("\n接口")
	fmt.Println(p.String())
	printString(p)
	printString(pe.addr)
	fmt.Println("值接收者和指针接收者")
	printString(&p)
	fmt.Println("工厂函数")
	p1 := NewPerson("cold")
	printString(p1)
	
	fmt.Println("组合")
	fmt.Println(pe.city)
	
	fmt.Println("类型断言")
	
	var s fmt.Stringer
	s = p1
	p2 := s.(*person)
	fmt.Println(p2)
	//a:=s.(address)
	//fmt.Println(a)
	a,ok := s.(address)
	if ok{
		fmt.Println(a)
	}else{
		fmt.Println("s不是一个address")
	}
}


type people struct{
	name string
	addr address
	// 组合
	address
}

type address struct{
	city string
}


func (p person) String() string{
	return fmt.Sprintf("the name is %s,age is %d",p.name,p.age)
}
// 以指针类型接收者实现接口的时候，只有对应的指针类型才被认为实现了该接口。
// func (p *person) String() string{
// 	return fmt.Sprintf("the name is %s,age is %d",p.name,p.age)
// }

func (addr address) String() string{
	return fmt.Sprintf("the address is %s",addr.city)
}
func printString(s fmt.Stringer){
	fmt.Println(s.String())
}

func NewPerson(name string) *person{
	return &person{name:name}
}