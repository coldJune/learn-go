package main

import(
	"fmt"
	"time"
	"sync"
)

func main(){
	fmt.Println("select timeout 模式")
	result := make(chan string)
	go func(){
		time.Sleep(8*time.Second)
		result <- "结果"
	}()
	
	select{
		case v:= <- result:
			fmt.Println(v)
		case <- time.After(5*time.Second):
			fmt.Println("超时了")
	}
	
	fmt.Println("\npipeline 模式")
	coms := buy(10)
	phones := build(coms)
	packs := pack(phones)
	// 输出测试，查看结果
	for p := range packs{
		fmt.Println(p)
	}
	fmt.Println("\n扇入扇出 模式")
	
	coms1 := buy(100)
	phones1 := build(coms1)
	phones2 := build(coms1)
	phones3 := build(coms1)
	
	allPhone := merge(phones1,phones2,phones3)
	packs1 := pack(allPhone)
	// 输出测试，查看结果
	for p := range packs1{
		fmt.Println(p)
	}
	fmt.Println("\nfutures 模式")
	
	vegetablesCh := washVegetables() //洗菜
	waterCh := boilWater()// 烧水
	fmt.Println("已经烧水洗菜，休息一会儿")
	time.Sleep(2*time.Second)
	fmt.Println("开始做菜")
	vegetables := <-vegetablesCh;
	fmt.Println(vegetables)
	water := <- waterCh;
	fmt.Println(water)
	fmt.Println("下锅",vegetables,water)
}

func buy(n int) <-chan string{
	out := make(chan string)
	go func(){
		defer close(out)
		for i := 1;i <= n; i++{
			out <- fmt.Sprint("配件",i)
		}
	}()
	return out
}

func build(in <- chan string) <-chan string{
	out := make(chan string)
	go func(){
		defer close(out)
		for c := range in{
			out <- "组装["+c+"]"
		}
	}()
	return out
}

func pack(in <- chan string) <-chan string{
	out := make(chan string)
	go func(){
		defer close(out)
		for c := range in{
			out <- "打包（"+c+"）"
		}
	}()
	return out
}

func merge(ins ... <-chan string) <-chan string{
	var wg sync.WaitGroup
	out := make(chan string)
	// 把一个channel中的数据发送到out中
	p:=func(in <-chan string){
		defer wg.Done()
		for c:= range in{
			out <- c
		}
	}
	wg.Add(len(ins))
	// 扇入，需要启动多个goroutine用于处于多个channel中的数据
	for _,cs:=range ins{
		go p(cs)
	}
	// 等待所有输入的数据ins处理完，再关闭输出out
	go func(){
		wg.Wait()
		close(out)
	}()
	return out
}

func washVegetables() <- chan string{
	vegetables := make(chan string)
	go func(){
		time.Sleep(5*time.Second)
		vegetables <- "洗好菜"
	}()
	return vegetables
}

func boilWater() <- chan string{
	water := make(chan string)
	go func(){
		time.Sleep(10*time.Second)
		water <- "烧好水"
	}()
	return water
}