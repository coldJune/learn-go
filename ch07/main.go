package main

import (
"fmt"
"strconv"
"errors"
)

func main(){
	fmt.Println("error 接口\n")
	i,err:=strconv.Atoi("a")
	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println(i)
	}
	
	fmt.Println("\nerror 工厂函数")
	sum,er := add(0,-1)
	if er != nil{
		fmt.Println(er)
	}else{
		fmt.Println(sum)
	}
	
	fmt.Println("\n自定义error")
	sumCe,cer := addThrowCommonError(0,-1)
	if cer != nil{
		fmt.Println(cer)
	}else{
		fmt.Println(sumCe)
	}
	
	fmt.Println("\nerror断言")
	sumAssert,cerAssert := addThrowCommonError(-1,2)
	if cm,ok:=cerAssert.(*commonError);ok{
		fmt.Println(cm.errorCode,cm.errorMsg)
	}else{
		fmt.Println(sumAssert)
	}
	
	fmt.Println("\n 错误嵌套")
	newError:=MyError{err,"数据问题"}
	fmt.Println(newError)
	
	e:=errors.New("原始错误e")
	w:=fmt.Errorf("wrap了一个错误：%w",e)//基于一个存在的error生成新的error
	fmt.Println(w)
	
	fmt.Println("\n errors.Unwrap 函数")
	fmt.Println(errors.Unwrap(w)) //用于获取被嵌套的 error
	
	fmt.Println(errors.Is(w,e)) //两个 error 相等或 err 包含 target 的情况下返回 true，其余返回 false
	
	fmt.Println("\n errors.As 函数")
	var com *commonError
	if errors.As(cer,&com){
		 fmt.Println("错误代码为:",com.errorCode,"，错误信息为：",com.errorMsg)
	}
}

func add(a, b int)(int, error){
	if a<0||b<0{
		return 0, errors.New("a或b不能为负数")
	}else{
		return a+b,nil
	}
}

type commonError struct{
	errorCode int // 错误码
	errorMsg string // 错误信息
}

func (ce *commonError) Error() string{
	return ce.errorMsg
}

func addThrowCommonError(a,b int)(int ,error){
	if a<0||b<0{
		return 0, &commonError{
			errorCode: 1,
			errorMsg: "a或b不能为负数",
		}
	}else{
		return a+b,nil
	}
}

type MyError struct{
	err error
	msg string
}

func (e *MyError) Error() string{
	return e.err.Error() +e.msg
}