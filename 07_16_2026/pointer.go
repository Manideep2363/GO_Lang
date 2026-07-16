package main

import "fmt"

func main() {
	x := 10
	p := &x
	*p = 100
	fmt.Println(x)
	fmt.Println(&x)
	fmt.Println(p)
	fmt.Println(*p)

	fmt.Println(x)
	fmt.Println(*p)
}
