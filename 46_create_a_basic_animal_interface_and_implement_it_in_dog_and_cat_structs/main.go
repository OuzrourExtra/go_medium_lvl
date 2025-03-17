package main

import (
	"fmt"
)

type Animal interface {
	Description()
	Barking()
	Walk()
}
type Cat struct {
	Size      string
	Sound     string
	Color     string
	WalkSound string
	Race      string
	Year      int
}

type Dog struct {
	Size      string
	Sound     string
	Color     string
	WalkSound string
	Race      string
	Year      int
}

func (c Cat) Barking() {
	fmt.Printf("\n %s I'm a CAT ! ", c.Sound)
}

func (d Dog) Barking() {
	fmt.Printf("\n %s I'm a DOG ! ", d.Sound)
}

func (c Cat) Walk() {
	fmt.Printf("\n When i move it sound like : %s !", c.WalkSound)
}

func (d Dog) Walk() {
	fmt.Printf("\n When i move it sound like : %s !", d.WalkSound)
}

func (c Cat) Description() {
	fmt.Printf("\n=================================\n")
	fmt.Printf("Hello , i'm a CAT ! i'm a %s cat . My Color is : %s , and my Race is : %s ... in fact , i have %d year ! Meow !", c.Size, c.Color, c.Race, c.Year)
	fmt.Printf("\n=================================")
}

func (d Dog) Description() {
	fmt.Printf("\n=================================\n")
	fmt.Printf("Hello , i'm a DOG ! i'm a %s dog . My Color is : %s , and my Race is : %s ... in fact , i have %d year ! Whof !", d.Size, d.Color, d.Race, d.Year)
	fmt.Printf("\n=================================")
}

func Description(animal Animal){
	animal.Description()
}

func Barking(animal Animal){
	animal.Barking()
}

func Walk(animal Animal){
	animal.Walk()
}

func main(){
	var animal Animal = nil

	animal = Cat{"big","Meow ! Meow !","Gray","Pew Pew!","Russian Blue",11}
	animal.Barking()
	animal.Walk()
	animal.Description()

	animal = Dog{"big","Whouf ! Whouf !","Gray","PAW PAW!","Germany Berger",1}
	animal.Barking()
	animal.Walk()
	animal.Description()
}