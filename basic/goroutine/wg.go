package main

import (
	"fmt"
	"net/http"
	"sync"
)

//1 .WaitGroup简介
//WaitGroup用于等待一组线程的结束。父线程调用Add方法来设定应等待的线程的数量。
//每个被等待的线程在结束时应调用Done方法。同时，主线程里可以调用Wait方法阻塞至所有线程结束。

//官网示例
func wg() {
	var wg sync.WaitGroup
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.somestupidname.com/",
		"http://www.baidu.com/",
	}
	for _, url := range urls {
		// Increment the WaitGroup counter.
		wg.Add(1)
		// Launch a goroutine to fetch the URL.
		go func(url string) {
			// Decrement the counter when the goroutine completes.
			defer wg.Done()
			// Fetch the URL.
			fmt.Println("请求的网站：", url)
			_, err := http.Get(url)
			fmt.Println(url, err)
		}(url)
	}
	// Wait for all HTTP fetches to complete.
	fmt.Println("阻塞中。。。")
	//阻塞至计数器减为0
	wg.Wait()
	fmt.Println("finished")
}

// Tip: waitGroup 不要进行copy
func errWg1() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//结构体为值类型，func内部wg与外部不是同一个 waitGroup了
		go func(i int, wg sync.WaitGroup) {
			fmt.Println(i)
			defer wg.Done()
		}(i, wg)
	}
	wg.Wait()
	fmt.Println("finished")
}

// Tip: waitGroup 的 Add 要在goroutine前执行
func errWg2() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		go func(i int) {
			wg.Add(1)
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("finished")
}

// Tip: waitGroup 的 Add 很大会有影响吗？
func errWg3() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(100)
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
			wg.Add(-100)
		}(i)
	}
	fmt.Println("finished")
}

func main() {
	errWg2()
}
