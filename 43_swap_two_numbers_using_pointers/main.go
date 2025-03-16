package main

import ("fmt")

func swap(number1 *int , number2 *int){
	fmt.Printf("=========================")
	fmt.Printf("\nBefore Swap : number1 = %d and number = %d",number1, number2)
	fmt.Printf(" \n ==> Memory Adress ( before ) : %v %v",&number1,&number2)
	*number1 , *number2 = *number2 , *number1
	fmt.Printf("\n=========================")
	fmt.Printf("\nAfter Swap : number1 = %d and number = %d",number1, number2)
	fmt.Printf(" \n ==> Memory Adress ( after ) : %v %v ( Same )",&number1,&number2)

}

func main(){
	number1 := 1
	number2 := 2
	swap(&number1,&number2)
	
}