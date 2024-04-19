package main

import (
	"fmt"
	"go-experimentation/utilz"
)

func main() {
	fmt.Println("hello world")
	var foo int
	var derp bool
	fmt.Println(foo)
	fmt.Println(derp)
	utilz.DoStuff()

	var avgTest int = utilz.Average(5, 10, 100)
	fmt.Println(avgTest)
}
