package main

import (
	"fmt"
	"strings"

	"github.com/ouzrourextra/go_medium_lvl/34_create_struct_of_a_person/err"
	"github.com/ouzrourextra/go_medium_lvl/34_create_struct_of_a_person/input"
)




type Person struct{
	name string
	age int
	email string
}

func verifyMailSyntax(str *string){
	if !strings.ContainsRune(*str,'@') || !strings.ContainsRune(*str,'.'){
		email , errEmail := input.InputString("(Wrong Syntaxe) Retype the e-mail of the person (string) ")
		err.Stop(errEmail)
		*str = email
		verifyMailSyntax(str)
	}
}

func fill (persons *[]Person, capacity int){
	for i:= 0 ; i<capacity ; i++{
		name , errName := input.InputString("Enter the name of the person (string) ")
		err.Stop(errName)
		age , errAge := input.InputNumber("Enter the age of the person (int) ")
		err.Stop(errAge)
		email , errEmail := input.InputString("Enter the e-mail of the person (string) ")
		err.Stop(errEmail)
		verifyMailSyntax(&email)
		
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