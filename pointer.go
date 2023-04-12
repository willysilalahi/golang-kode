package main

import "fmt"

func main() {
	var x int = 1
	var y int
	var ip *int

	ip = &x
	y = *ip

	fmt.Println(ip)
	fmt.Println(y)
	fmt.Println("---------------")

	//new
	ptr := new(int)
	*ptr = 14
	fmt.Println(ptr)
	fmt.Println(*ptr)

	test := new(int)
	fmt.Println(test)
}
