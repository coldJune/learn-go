package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main(){
	fmt.Println("动态扩容")
	ss := []string{"cold","june"}
	fmt.Println("切片ss长度为",len(ss),"容量为",cap(ss))
	fmt.Println(ss)
	ss = append(ss,"yes","no")
	fmt.Println("切片ss长度为",len(ss),"容量为",cap(ss))
	fmt.Println(ss)
	
	fmt.Println("\n数据结构")
	a1 := [2]string{"cold","june"}
	s1 := a1[0:1]
	s2 := a1[:]
	//打印出s1和s2的Data值是一样的
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s2)).Data)
	sh1 := (*slice)(unsafe.Pointer(&s1))
	fmt.Println(sh1.D,sh1.L,sh1.C)
	
	fmt.Println("\n高效的原因")
	// 切片在赋值传参时并不会把所有元素复制一遍，而只是复制SliceHeader的三个字段，底层依然共用一个数组
	fmt.Printf("函数main数组指针:%p\n",&a1)
	arrayF(a1)
	fmt.Println((*reflect.SliceHeader)(unsafe.Pointer(&s1)).Data)
	sliceF(s1)
	
	fmt.Println("\nstring和[]byte互转")
	s := "cold"
	fmt.Printf("s的内存地址：%d\n",(*reflect.StringHeader)(unsafe.Pointer(&s)).Data)
	b := []byte(s)
	fmt.Printf("b的内存地址：%d\n",(*reflect.StringHeader)(unsafe.Pointer(&b)).Data)
	s3 := string(b)
	fmt.Printf("s3的内存地址：%d\n",(*reflect.StringHeader)(unsafe.Pointer(&s3)).Data)
	fmt.Println(s,string(b),s3)
	
	s4 := *(*string)(unsafe.Pointer(&b))
	fmt.Println(s4)
	
}

func arrayF(a [2]string){
	fmt.Printf("函数arrayF数组指针:%p\n",&a)
}

func sliceF(s []string){
	fmt.Printf("函数sliceF Data:%d\n",(*reflect.SliceHeader)(unsafe.Pointer(&s)).Data)

}
type slice struct{
	D uintptr
	L int
	C int
}