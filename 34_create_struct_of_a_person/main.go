package main

import (
	"github.com/ouzrourextra/go_medium_lvl/34_create_struct_of_a_person/input"
	"github.com/ouzrourextra/go_medium_lvl/34_create_struct_of_a_person/err"
	"fmt"
)




type Person struct{
	name string
	age int
	email string
}

func fill (persons *[]Person, capacity int){
	for i:= 0 ; i<capacity ; i++{
		name , errName := input.InputString("Enter the name of the person (string) ")
		err.Stop(errName)
		age , errAge := input.InputNumber("Enter the age of the person (int) ")
		err.Stop(errAge)
		email , errEmail := input.InputString("Enter the e-mail of the person (string) ")
		err.Stop(errEmail)
		
		*persons = append(*persons, Person{name: name, age: int(age), email: email})

	}
}

func main(){
	capacity , errCapactiy := input.InputNumber("Enter the number of Persons")
	err.Stop(errCapactiy)
	ListOfPersons := make([]Person,0,capacity)
	fmt.Printf("Before : %v\n",ListOfPersons)
	fill(&ListOfPersons,int(capacity))
	fmt.Printf("\nAfter : %v",ListOfPersons)
}