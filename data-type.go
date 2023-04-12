package main

import "fmt"

func main() {
	var x int32 = 1
	var y int16 = 1
	y = int32(y)
	fmt.Println(x == y)
	fmt.Println(y)
}
