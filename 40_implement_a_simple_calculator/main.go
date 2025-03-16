package main

import (
	"fmt"
	"github.com/ouzrourextra/go_medium_lvl/40_implement_a_simple_calculator/input"
)

func calculator(nb1 float64 , nb2 float64 , operation *string){
	switch *operation{
	case "+":
		fmt.Println("===============================")
		fmt.Printf(" ==> The Result is : %.2f + %.2f = %.2f \n",nb1,nb2,nb1+nb2)
	case "-":
		fmt.Println("===============================")
		fmt.Printf(" ==> The Result is : %.2f - %.2f = %.2f \n",nb1,nb2,nb1-nb2)
	case "*":
		fmt.Println("===============================")
		fmt.Printf(" ==> The Result is : %.2f * %.2f = %.2f \n",nb1,nb2,nb1*nb2)
	case "/":
		fmt.Println("===============================")
		fmt.Printf(" ==> The Result is : %.2f / %.2f = %.2f \n",nb1,nb2,nb1/nb2)
	default :
		*operation  = input.InputString("(Wrong) Re-type the operation")
		calculator(nb1,nb2,operation)
	}
}
func main(){
	fmt.Println("===============================")
	fmt.Println("Calculator v1.0")
	for{
	number1 := input.InputNumber("Enter Number 1")
	number2 := input.InputNumber("Enter Number 2")
	operation := input.InputString("Enter the operation ( + , - , * , / )")
	calculator(number1,number2,&operation)
	}
	
}