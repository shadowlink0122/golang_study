package main

import "fmt"

func main() {
	var a int = 10
	var pa *int = &a

	fmt.Println(pa)
	fmt.Println(*pa)
}
