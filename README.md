# practice-go

practice for go

## 算法
- [合并两个有序链表](./demo001/main.go)
- [有序数组,原地删除重复项,返回新数组](./demo002/main.go)
- [递归求和](./demo003/main.go)
- [斐波那契数列求和(黄金分割数列、兔子数列)](./demo004/main.go)

## 数组排序
- [冒泡排序](./demo006/main.go)
- [选择排序](./demo006/main.go)
- [插入排序](./demo006/main.go)

## 链表排序
- [冒泡排序](./demo007/main.go)
- [插入排序](./demo007/main.go)

## EncryptAes ECB CBC CFB
- [AES加密(ECB/CBC/CFB)](./utils/crypted_aes.go)
## channel
```golang
// a simple producer customer
package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int, 10)
	go func() {
		for {
			c <- 1
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	select {}
}
```

# questions
- [goroutine、runtime.Gosched()](/question/coding.md)