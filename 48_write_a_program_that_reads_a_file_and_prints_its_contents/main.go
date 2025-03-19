package main

import (
	"fmt"
	"os"
)
func writeFile(name string){
	_ , err := os.Stat(name) 
	if err == nil {
		reader , _ := os.ReadFile(name)
		fmt.Println("FROM '",name,"' : ",string(reader))
	}else{
		os.WriteFile(name,[]byte("Hello From The new file , My Name is : ILYAS OUZROUR !"),0644)
		reader , _ := os.ReadFile(name)
		fmt.Println("FROM '",name,"' : ",string(reader))
	}
	if name == "delete-it.txt"{
		os.Remove(name)
	}
}
func main(){
	writeFile("Text.txt")
	writeFile("Text2.txt")
	writeFile("delete-it.txt")
}