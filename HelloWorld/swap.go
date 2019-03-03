package main

import "fmt"

func swap(a int, b int) (int, int){
	return b, a
}

func main(){
	fmt.Println(swap(2, 5))

	f := func(a int, b int) (int, int){
		return b, a
	}

	fmt.Println(f(2, 5))
}
