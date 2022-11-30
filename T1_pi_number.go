package main

import (
	"fmt"
)

func mainT1() {
	var π = 22 / 7.0
	fmt.Println(π)

	/*
		//not compiling: int can not divide by float
		a1, b1 := 22, 7.0
		var π1 = a1 / b1
		fmt.Println(π1)
	*/
}
