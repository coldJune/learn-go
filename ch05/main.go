package main

import "fmt"
import "errors"

func main(){
	fmt.Println("函数")
	result := sum(1,2)
	fmt.Println(result)
	
	result1,err := sum1(-1,2)
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(result1)
	}
	
	result2, err2 := sumNamed(1,-1)
	if err2 != nil{
		fmt.Println(err2)
	}else{
		fmt.Println(result2)
	}
	
	fmt.Println("sum(1,2)=",sumChangeParams(1,2))
	fmt.Println("sum(1,2,4)=",sumChangeParams(1,2,4))
	
	fmt.Println("匿名函数")
	sum2 := func(a int,b int)int{
		return a+b
	}
	fmt.Println(sum2(1,2))
	
	fmt.Println("\n函数闭包")
	cl := colsure()
	
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())
	fmt.Println(cl())
	
	fmt.Println("\n方法")
	age:=Age(15)
	age.add("岁")
	
	(&age).modify()
	age.add("岁")
	age = Age(20)
	age.add("岁")
	age.modify()
	age.add("岁")
}

func sum(a int,b int) int{
	return a+b
}

func sum1(a int, b int)(int,error){
	if a<0||b<0{
		return 0,errors.New("a或b不能小于0")
	}
	return a+b,nil
}

func sumNamed(a int,b int)(sum int,err error){
	if a<0||b<0{
		return 0,errors.New("a或b不能小于0")
	}
	sum = a+b
	err = nil
	return
}

func sumChangeParams(params ...int)int{
	sum :=0 
	for _,i:=range params{
		sum+=i
	}
	return sum
}

func colsure() func() int{
	i := 0 
	return func() int {
		i++
		return i
	}
}

type Age uint
func (age Age) add(a string){
	fmt.Println("年龄为：",age,a)
}

func (age *Age) modify(){
	*age = Age(30)
}