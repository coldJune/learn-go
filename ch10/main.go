package main

import (
"fmt"
"sync"
"time"
"context"
)

func main(){
	fmt.Println("使用管道控制")
	var wg sync.WaitGroup
	wg.Add(1)
	stopCh := make(chan bool) //用来停止监控
	go func(){
		defer wg.Done()
		watchDog(stopCh,"监控1")
	}()
	time.Sleep(5*time.Second) //先监控5s
	stopCh<-true
	wg.Wait()
	
	fmt.Println("使用context控制")
	var wgContext sync.WaitGroup
	wgContext.Add(4)
	ctx,stop := context.WithCancel(context.Background())
	go func(){
		defer wgContext.Done()
		watchDogContext(ctx,"context监控1")
	}()
	go func(){
		defer wgContext.Done()
		watchDogContext(ctx,"context监控2")
	}()
	go func(){
		defer wgContext.Done()
		watchDogContext(ctx,"context监控3")
	}()
	
	valCtx:=context.WithValue(ctx,"userId",2)
	go func(){
		defer wgContext.Done()
		getUser(valCtx)
	}()
	time.Sleep(5*time.Second) //监控5s
	stop()
	wgContext.Wait()
}

func watchDog(stopCh chan bool,name string){
	// 开启for select 循环，一直后台监控
	for{
		select{
			case <- stopCh:
				fmt.Println(name,"停止监控")
				return
			default:
				fmt.Println(name,"正在监控")
		}
		time.Sleep(time.Second)
	}
}

func watchDogContext(ctx context.Context, name string){
	// 开启for select循环，一直监控
	for{
		select{
			case <- ctx.Done():
				fmt.Println(name,"停止监控")
				return
			default:
				fmt.Println(name,"正在监控")
		}
		time.Sleep(1* time.Second)
	}
}

func getUser(ctx context.Context){
// 开启for select循环，一直监控
	for{
		select{
			case <- ctx.Done():
				fmt.Println("获取用户","退出")
				return
			default:
				userId := ctx.Value("userId")
				fmt.Println("获取用户","用户Id为：",userId)
				time.Sleep(1*time.Second)
		}
		time.Sleep(1* time.Second)
	}
}