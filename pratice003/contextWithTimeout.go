package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 多个goroutine 并发执行 设置超时时间，也允许提前结束，允许失败，不在意返回结果
func main() {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()

		down := make(chan int)

		go func(ctx context.Context) {
			time.Sleep(time.Second * 5)

			down <- 1
		}(ctx)

		select {
		case <-down:
			fmt.Println("child process success!", time.Now())
			return
		case <-ctx.Done():
			fmt.Println("child process exit!", time.Now())
			return
		}
	}(ctx)

	fmt.Println("main goroutine start" + time.Now().Format("2006-01-02 15:04:05"))
	wg.Wait()
	fmt.Println("main goroutine end" + time.Now().Format("2006-01-02 15:04:05"))

	time.Sleep(time.Second * 30)
}
