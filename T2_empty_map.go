package main

import (
	"fmt"
)

func main() {
	var myMap map[string]int
	fmt.Println(myMap["g"])

	v := myMap["h"]
	fmt.Println(v)
}
