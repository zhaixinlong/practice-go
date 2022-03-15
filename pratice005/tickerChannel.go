package main

import (
	"fmt"
	"sync"
	"time"
)

func producer(ch chan int) {
	ticker := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-ticker.C:
			ch <- 1
			fmt.Printf("p %d ch %d \n", time.Now().Second(), len(ch))
		}
	}
}

func consumer(ch chan int) {
	ticker := time.NewTicker(time.Second * 20)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("consumer %d ch %d \n", time.Now().Second(), len(ch))
			fmt.Printf("consumer %d  v %d \n", time.Now().Second(), <-ch)
		}
	}
}

func main() {
	// 5个生产者，每10秒生产一个struct
	// 10个消费者，每20秒消费一个struct
	// 缓冲区 20

	ch := make(chan int, 20)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			producer(ch)
		}()
		fmt.Printf("wg producer 1 5 %d ch %d \n", time.Now().Second(), len(ch))
	}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			consumer(ch)
		}()
		fmt.Printf("wg consumer 1 10 %d ch %d \n", time.Now().Second(), len(ch))
	}

	wg.Wait()
}
