package main

import "fmt"

func main(){
	var a [5]int

	a[2] = 2
	a[4] = 5

	fmt.Println(a)

	b := [3]int{1, 2, 3}
	fmt.Println(b)

	c := [...]int{1, 2, 3, 4, 5}
	fmt.Println(c)
}