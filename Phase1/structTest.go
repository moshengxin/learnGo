package main

import (
	"fmt"
	"time"
)

type User1 struct {
	name string
	age  int
}

func main() {

	u := User1{}
	u.name = "zm"
	u.age = 111
	fmt.Println(u, time.Now())
}
