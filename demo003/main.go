package main

import "fmt"

// 递归求和
// 1 + 2 + 3 + ..... + n-1 + n

func calcSum(num int) int {
	if num == 1 {
		return 1
	}
	return num + calcSum(num-1)
}

func main() {
	s := calcSum(10)
	fmt.Printf("sum s %d \n", s)

	s = calcSum(100)
	fmt.Printf("sum s %d \n", s)
}
