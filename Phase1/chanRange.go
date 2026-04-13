package main

import "fmt"

func main() {

	quee1 := make(chan string, 3)
	quee1 <- "1"
	quee1 <- "2"
	quee1 <- "3"
	close(quee1)

	var s string
	for s = range quee1 {
		fmt.Println(s)
	}

}
