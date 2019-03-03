package main

import "fmt"

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6}

	b := a[:]
	c := a[2:]
	d := a[:4]
	e := a[2:4]

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)

	fmt.Println(len(a))

	fmt.Println(cap(a))
}
