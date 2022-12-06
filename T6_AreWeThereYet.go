package main

import (
	"fmt"
	"time"
)

func mainT6() {
	timeout := 3
	fmt.Printf("before ")
	//Not compiled!!!
	//time.Sleep(timeout * time.Millisecond)
	time.Sleep(3 * time.Millisecond)
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	fmt.Println("after")
}
