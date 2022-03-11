package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// many goroutines access the same variable slice , make consistent

	// not consistent
	notConsistent()

	// not really consistent
	useChannelLooksConsistentButNot()

	// consistent
	useMutexConsistent()
}

func notConsistent() {
	var is []int
	var wg sync.WaitGroup

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			is = append(is, i)
		}
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			is = append(is, i)
		}
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			is = append(is, i)
		}
	}(&wg)

	wg.Wait()
	fmt.Printf("is %d \n", len(is))
}

func useChannelLooksConsistentButNot() {
	var is []int
	v := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			v <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			v <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			v <- i
		}
	}()

	go func() {
		for vv := range v {
			is = append(is, vv)
		}
	}()

	time.Sleep(time.Millisecond * 100)
	fmt.Printf("is %d \n", len(is))
}

func useMutexConsistent() {
	var is []int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			mu.Lock()
			is = append(is, i)
			mu.Unlock()
		}
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			mu.Lock()
			is = append(is, i)
			mu.Unlock()
		}
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			mu.Lock()
			is = append(is, i)
			mu.Unlock()
		}
	}(&wg)

	wg.Wait()
	fmt.Printf("is %d \n", len(is))
}
