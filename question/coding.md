#  questions

## coding
1、goroutine调度，runtime.Gosched()作用。以下代码输出结果
```golang
package main

import (
	"fmt"
	"runtime"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}

func main() {
	go say("world") //开一个新的Goroutines执行
	say("hello")    //当前Goroutines执行

	// 不一定输出什么，大部分输出五个hello，应该是取决于goroutine调度？
	// time.Sleep(time.Second * 2) 加sleep不一样的结果
}
```