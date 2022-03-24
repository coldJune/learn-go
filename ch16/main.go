package main

import (
	"fmt"
	"unsafe"
)

func main(){
	fmt.Println("\n指针类型转换")
	i := 10
	ip := &i
	// var fp *float64 = (*float64)(ip) 不同的指针类型不能强制转换
	var fp *float64 = (*float64) (unsafe.Pointer(ip)) // unsafe.Pointer作为不同指针类型之间转换的桥梁
	fmt.Println(fp)
	
	fmt.Println("\nuintptr 指针类型")
	p := new(person)
	// Name是person的第一个字段不用偏移，即可通过指针修改
	pName := (*string)(unsafe.Pointer(p))
	*pName = "jun"
	// Age并不是person的第一个字段，所以需要进行偏移，才能正确定位到Age字段的内存，才能修改
	pAge := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(p))+unsafe.Offsetof(p.Age)))
	*pAge = 20
	fmt.Println(*p)
	
	fmt.Println("\nunsage.Sizeof")
	fmt.Println(unsafe.Sizeof(true))
	fmt.Println(unsafe.Sizeof(int8(10)))
	fmt.Println(unsafe.Sizeof(int16(100)))
	fmt.Println(unsafe.Sizeof(int32(1000)))
	fmt.Println(unsafe.Sizeof(int64(10000)))
	fmt.Println(unsafe.Sizeof(int(100000)))
	fmt.Println(unsafe.Sizeof(string("1")))
	fmt.Println(unsafe.Sizeof([]string{"1","2"}))

	
}

type person struct {

   Name string

   Age int

}