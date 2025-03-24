package main

import (
	"os"
	"log"
	"fmt"
)


func exist(filename string){
	dataBefore , errReading := os.ReadFile(filename)
	if errReading!=nil{
		log.Fatal(errReading)
	}
	fmt.Printf("\nData : %v",dataBefore)
}

func main(){
	exist("test.txt")
}