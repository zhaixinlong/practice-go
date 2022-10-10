package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var mu sync.RWMutex
	m := map[string]string{
		"k": "v",
	}

	go func(mu *sync.RWMutex) {
		time.Sleep(time.Second * 1)
		for {
			fmt.Printf("w-1 \n")

			mu.Lock()
			fmt.Printf("w-11 \n")
			time.Sleep(time.Second * 5)
			m["k"] = "vv"

			time.Sleep(time.Second * 5)
			mu.Unlock()
		}
	}(&mu)

	go func(mu *sync.RWMutex) {
		for {
			fmt.Printf("r-1\n")

			mu.RLock()
			fmt.Printf("m-1: %v\n", m["k"])
			time.Sleep(time.Second * 1)
			fmt.Printf("m-1: %v\n", m["k"])
			mu.RUnlock()
		}
	}(&mu)

	select {}
}
