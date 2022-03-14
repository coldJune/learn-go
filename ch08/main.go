package main

import (
"fmt"
"time"
)
func main(){
	go fmt.Println("jun")
	fmt.Println("main")
	time.Sleep(time.Second)
	
	fmt.Println("\n通道channel")
	ch := make(chan string)
	go func(){
		fmt.Println("cold")
		ch <- "go 完成"
	}()
	
	fmt.Println("main")
	v := <-ch
	fmt.Println("接受chan的:",v)
	
	fmt.Println("\n有缓冲通道channel")
	cacheCh := make(chan int,5)
	cacheCh <- 2
	cacheCh <- 3
	fmt.Println("cacheCh的容量为:",cap(cacheCh),",元素个数为：",len(cacheCh))
	
	fmt.Println("\n单向channel")
	 //onlySend := make(chan<-int)向Channel写
	 //onlyReceive := make(<- chan int)从channel读
	go send(cacheCh)
	go recv(cacheCh)
	time.Sleep(time.Second)
	close(cacheCh)
	
	firstCh := make(chan string)
	sencondCh := make(chan string)
	threeCh := make(chan string)
	go func(){
		firstCh <- downloadFile("firstCh")
	}()
	go func(){
		sencondCh <- downloadFile("sencondCh")
	}()
	go func(){
		threeCh <- downloadFile("threeCh")
	}()
	
	// 开始select多路复用，哪个channel能获取到值就说明哪个最先好
	select{
		case filePath := <- firstCh:
			fmt.Println(filePath)
		case filePath := <- sencondCh:
			fmt.Println(filePath)
		case filePath := <- threeCh:
			fmt.Println(filePath)
	}
}

// 只能向chan写数据
func send(out chan <- int){
	for i:=0; i<10;i++{
		out <- i
	}
}
// 只能从chan读数据
func recv(in <- chan int){
	for i := range in{
		fmt.Println(i)
	}
}

func downloadFile(chanName string) string{
	time.Sleep(time.Second)
	return chanName+":filePath"
}

