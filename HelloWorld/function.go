package main

import "fmt"

func sayHello(){
	fmt.Println("Hello")
}

func sayName(name string){
	fmt.Println(name)
}

func getMessage(name string) string{
	msg := "My name is " + name
	return msg
}

func getHelloMessage(name string) (msg string){
	msg = "Hello " + name
	return 
}

func main(){
	sayHello()
	sayName("Takashi")
	fmt.Println(getMessage("Takashi"))
	fmt.Println(getHelloMessage("Takashi"))
}
