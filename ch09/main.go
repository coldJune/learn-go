package main

import (
	"fmt"
	"time"
	"sync"
)

var sum = 0

var (
	sumMute int
	mutex sync.Mutex
)
func main(){
	for i:=0;i<1000;i++{
		go add(10)
	}
	time.Sleep(2*time.Second)
	fmt.Println("sum:",sum)
	
	fmt.Println("\n同步原语")
	for i:=0;i<1000;i++{
		go addMutex(10)
	}
	time.Sleep(2*time.Second)
	fmt.Println("sumMute:",sumMute)
	
	fmt.Println("\n读写锁")
	sumMute = 0
	for i:=0;i<1000;i++{
		go addMutex(10)
	}
	for i:=0;i<10;i++{
		go fmt.Println("和为:",readSum())
	}
	time.Sleep(2*time.Second)
	fmt.Println("\nWaitGroup")
	run()
	fmt.Println("\nOnce")
	doOnce()
	fmt.Println("\nCond")
	race()
}

func add(i int){
	sum += i
}

func addMutex(i int){
	mutex.Lock()
	defer mutex.Unlock()
	sumMute += i
}
var mutexRW sync.RWMutex

func readSum() int{
	mutexRW.RLock()
	defer mutexRW.RUnlock()
	b := sum
	return b
}

func run(){
	var wg sync.WaitGroup
	wg.Add(110)
	for i:=0;i<100;i++{
		go func(){
			defer wg.Done()
			add(10)
		}()
	}
	
	for i:=0;i<10;i++{
		go func(){
			defer wg.Done()
			fmt.Println("和为:",readSum())
		}()
	}
	wg.Wait()
}

func doOnce(){
	var once sync.Once
	
	onceBody := func(){
		fmt.Println("only once")
	
	}
	
	done := make(chan bool)
	for i:=0;i<10;i++{
		go func(){
			once.Do(onceBody)
			done <-true
		}()
	}
	for i:=0;i<10;i++{
		<-done
	}
}

func race(){
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	wg.Add(11)
	for i:=0;i<10;i++{
		go func(num int){
			defer wg.Done()
			fmt.Println(num,"已就位")
			cond.L.Lock()
			cond.Wait()
			fmt.Println(num,"号开始跑")
			cond.L.Unlock()
		}(i)
	}
	time.Sleep(2*time.Second)
	go func(){
		defer wg.Done()
		fmt.Println("裁判已就位，准备发令枪")
		fmt.Println("开始~~")
		cond.Broadcast()
	}()
	
	wg.Wait()
}