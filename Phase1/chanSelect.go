package main

import "fmt"

func main() {

	a1 := make(chan string)
	a2 := make(chan string)

	// 匿名协程1
	go func() {
		a1 <- "a1"
	}()

	// 匿名协程2
	go func() {
		a2 <- "a2"
	}()

	// 并发执行，打印顺序 无保证
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-a1:
			fmt.Println("received", msg1)
		case msg2 := <-a2:
			fmt.Println("received", msg2)
		}
	}

}
