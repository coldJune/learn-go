package main

import (
	"fmt"
)

func main(){
	name := "cold"
	nameP := &name
	fmt.Println("name的值为：",name)
	fmt.Println("name的地址为：",nameP)
	
	nameV := *nameP
	fmt.Println("nameP指向的值为:",nameV)
	
	*nameP = "june" //修改指针指向的值
	fmt.Println("*nameP指向的值为：",*nameP)
	fmt.Println("name的值为：",name)
	fmt.Println("\n指针参数")
	
	age := 26
	modifyAge(age)
	fmt.Println("age的值为:",age)
	modifyAgeRef(&age)
	fmt.Println("age的值为:",age)

}
func modifyAge(age int){
	age = 20
}

func modifyAgeRef(age *int){
	*age = 20
}