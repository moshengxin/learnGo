package main

import (
	"fmt"
	"time"
)

func main() {
	// main主线程不能提前结束，否则定时器将不能执行
	time.AfterFunc(2*time.Second, func() {
		fmt.Println("2秒后，执行方法")
	})
	timer1 := time.NewTimer(2 * time.Second)

	<-timer1.C
	fmt.Println("Timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("Timer 2 fired")
	}()
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("Timer 2 stopped")
	}

	time.Sleep(2 * time.Second)

}
