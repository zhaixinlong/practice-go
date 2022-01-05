package main

import "fmt"

// 斐波那契数列
// 1、1、2、3、5、8、13、21、34、……
func calcFeibonaqie(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return calcFeibonaqie(n-1) + calcFeibonaqie(n-2)
}

func main() {
	t := calcFeibonaqie(24)
	fmt.Printf("sum t %d \n", t)
}
