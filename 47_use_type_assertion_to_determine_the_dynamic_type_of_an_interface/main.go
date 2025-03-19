package main

import "fmt"

func detectTheType(data interface{}){
	switch v:= data.(type){
	case string : 
		fmt.Printf("%q is a string.\n",v)
	case int : 
		fmt.Printf("%d is a int.\n",v)
	case bool :
		fmt.Printf("%t is a boolean.\n",v)
	case rune : 
		fmt.Printf("%c is a rune.\n",v)
	default :
		fmt.Printf("Unknown Type !\n")
	}
}

func main(){
	var number int = 2
	var str string = "Hello"
	var char rune = '@'
	var boolean bool = true
	var float float64 = 3.14
	detectTheType(number)
	detectTheType(str)
	detectTheType(char)
	detectTheType(boolean)
	detectTheType(float)
}