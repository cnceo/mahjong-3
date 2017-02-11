package main

import (
	"fmt"
	"time"
)

func main() {
	a := []int{1, 2, 3, 4}
	fmt.Println(a[0:2])
	fmt.Println("hello world")

	notify := make(chan int, 1)
	start := time.Now()
	go func() {
		time.Sleep(time.Microsecond*5)
		notify <- 3
	}()

	breakTimerTime := time.Duration(0)
	timeout := time.Second * 10
	for  {
		timer := timeout - breakTimerTime
		fmt.Println("timer :", timer, timeout, breakTimerTime)
		select {
		case <- time.After(timer):
			fmt.Println(time.Now().Sub(start))
			return
		case <- notify:
			breakTimerTime += time.Now().Sub(start)
			fmt.Println("read notify", breakTimerTime)
		}
	}
}