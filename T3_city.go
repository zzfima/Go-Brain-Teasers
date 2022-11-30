package main

import (
	"fmt"
	"unicode/utf8"
)

func mainT3() {
	city1 := "Kraków"
	fmt.Println(len(city1))

	city2 := "Krakow"
	fmt.Println(len(city2))

	c1 := 'o'
	c2 := 'ó'
	c3 := 'ж'
	fmt.Println(c1, c2, c3)
	fmt.Printf("len of o, ó, ж %d %d %d\n", len("o"), len("ó"), len("ж"))

	fmt.Printf("%s includes %d runes", city1, utf8.RuneCountInString(city1))
}
