package main

import (
	"fmt"
)

func PointerModify(modified *int){
	*modified = 1996
}

func main(){
	number := 9
	fmt.Printf("\n Before : %d",number)
	PointerModify(&number)	
	fmt.Printf("\n After : %d",number)
}