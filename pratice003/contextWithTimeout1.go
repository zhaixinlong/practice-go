package main

import (
	"context"
	"fmt"
	"time"
	// ====== 可以引用Golang原生语言包 ====== //
)

// IWorkload 请勿修改接口
type IWorkload interface {
	// Process内包含一些耗时的处理，可能是密集计算或者外部IO
	Process()
}

var TimeoutError = fmt.Errorf("timeout")

// 问题1：请编写函数Question1的实现如下功能
// 该函数输入一个IWorkload实例，请调用其Process函数一次，
// 调用完毕则Question1返回，此时返回的error应为空
// 当Process函数执行5秒仍未能结束时，让Question1函数不再等待
// 立即返回TimeoutError
//
// 注意：题目要求只调用Process一次
// 注意：超时时间固定5秒，请不要修改Question1函数的输入、输出定义
// 提示：请尽量使用规范的代码风格，使代码整洁易读
// 提示：如果也实现了测试代码，请一并提交，将有利于分数评定

func Question1(workload IWorkload) (err error) {
	// ====== 在这里书写代码 ====== //
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	ch := make(chan bool)

	go func(workload IWorkload, ch chan bool) {
		fmt.Printf("run Process \n")

		workload.Process()
		ch <- true
	}(workload, ch)

	select {
	case <-ctx.Done():
		fmt.Printf("time out %v \n", ctx.Err())
		return TimeoutError
	case <-ch:
		fmt.Printf("ch \n")
	}
	return nil
}

type WorkloadImpl struct {
}

func (wl *WorkloadImpl) Process() {
	time.Sleep(10 * time.Second)
}

func main() {
	err := Question1(new(WorkloadImpl))
	fmt.Printf("err %v \n", err)
}
