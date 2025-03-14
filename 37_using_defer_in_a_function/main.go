package main

import ("fmt")

func sayHello(){
	// Runned in LIFO mod ( Last-in First-Out )
	defer fmt.Printf(" 3 ")
	defer fmt.Printf(" 2 ")
	defer fmt.Printf(" 1 ")
	defer fmt.Printf("The Function run is finished but , i'm gonna count to 3 : ")
	// The core of the function
	fmt.Printf("Hello world !\n")
}

func main(){
	sayHello()
}