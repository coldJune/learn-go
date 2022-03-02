package main

import "fmt"
import "unicode/utf8"

func main(){
	fmt.Println("数组")
	array:=[2]string{"a","b"}
	fmt.Println(array[1])
	
	array1:=[...]string{"1","2","3"}
	fmt.Println(array1[2])
	
	array2:=[5]string{1:"b",3:"3"}
	fmt.Println(array2[2],array2[3])// 没有初始化的索引默认是对应类型的零值
	
	for i:=0;i<len(array);i++{
		fmt.Printf("数组索引:%d,对应值:%s\n",i,array[i])
	}
	
	fmt.Println("\n数组循环")
	for i,v:=range array{
		fmt.Printf("数组索引:%d,对应值:%s\n",i,v)
	}
	
	for _,v:=range array{
		fmt.Printf("对应值:%s\n",v)
	}
	
	fmt.Println("\n切片")
	array3 := [5]string{"1","2","3","4","5"}
	slice := array3[2:5]
	fmt.Println(slice)
	slice[1] = "a"
	fmt.Println("slice:",slice)
	fmt.Println("array3:",array3)
	
	fmt.Println("\n切片声明")
	slice1 := make([]string,4,8)
	fmt.Println(slice1)
	slice2 := []string{"a","b"}
	fmt.Println(len(slice2),cap(slice2))
	
	fmt.Println("\nappend")
	slice3:=append(slice1,"f")
	fmt.Println(slice3)
	slice3 = append(slice1,"g","g")
	fmt.Println(slice3)
	slice3 = append(slice1,slice2...)
	fmt.Println(slice3)
	for i,v:=range slice3{
		fmt.Printf("切片索引:%d,对应值:%s\n",i,v)
	}
	
	fmt.Println("\nMap")
	nameAgeMap:=make(map[string]int)
	nameAgeMap["jun"] = 26
	fmt.Println(nameAgeMap["jun"])
	fmt.Println(nameAgeMap)
	age, ok := nameAgeMap["jun"]
	if ok{
		fmt.Println(age)
	}
	delete(nameAgeMap,"jun")
	fmt.Println(nameAgeMap)
	nameAgeMap["cold"] = 20
	nameAgeMap["jun"] = 26
	for k,v:=range nameAgeMap{
		fmt.Printf("map键:%s,对应值:%d\n",k,v)
	}
	fmt.Println(len(nameAgeMap))
	
	fmt.Println("\nString 和[]byte")
	s := "cold 俊"
	bs := []byte(s)
	fmt.Println("s:",s,"bs:",bs)
	fmt.Println("s[3]",s[3],"bs[7]:",bs[7])
	fmt.Println("len(s):",len(s),"len(bs):",len(bs)) // len计算字节
	fmt.Println(utf8.RuneCountInString(s))
	
	for i,r:=range s{ //自动隐式解码Unicode字符串
		fmt.Println(i,r)
	}
}