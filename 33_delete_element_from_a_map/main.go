package main

import (
	"fmt"
	"bufio"
	"log"
	"strconv"
	"strings"
	"os"
)

/*
inputNumber is a function for inputing a number
*/
func inputNumber(sentence string) (nb int64, err error) {
	fmt.Println("=======================")
	fmt.Printf("%s: ",sentence)
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	nb, err1 := strconv.ParseInt(str, 10, 64)
	if err1 != nil {
		return 0, err1
	} else {
		return nb, nil
	}
}



func fillTheMap(MyMap *map[string]int,capacity int){
	fmt.Printf("\n=========================")

	for i:= 0 ; i<capacity ; i++ {
		fmt.Printf("\nIndex of Element %d (string) : ",i+1)
		readerString := bufio.NewReader(os.Stdin)
		String , errString := readerString.ReadString('\n')
		if errString != nil {
			log.Fatal(errString)
		}
		String = strings.TrimSpace(String)
		fmt.Printf("Value of Element %d (int) : ",i+1)

		readerInt := bufio.NewReader(os.Stdin)
		IntString , errInt := readerInt.ReadString('\n')
		if errInt != nil{
			log.Fatal(errInt)
		}
		Int , errConv := strconv.Atoi(strings.TrimSpace(IntString))
		if errConv != nil{
			log.Fatal(errConv)
		}
		(*MyMap)[String] = Int
		fmt.Printf("INFO : Added Successfully !\n")
		fmt.Printf("=========================\n")

	}
}



func searchIndex(myMap map[string]int,word string)bool{
	exist:= false
	for str , _ := range myMap{
		if str == strings.TrimSpace(word) {
			exist = true
		} 
	}
	return exist
}



func main(){
	numberOfElements , err := inputNumber("Enter the number of Elements in the map")
	if err != nil{
		log.Fatal(err)
	}
	mapCreated := make(map[string]int)

	fillTheMap(&mapCreated,int(numberOfElements))
	fmt.Printf("Enter the key you want to delete if it exist : ")
	reader := bufio.NewReader(os.Stdin)
	word , err := reader.ReadString('\n')
	word = strings.TrimSpace(word)
	if err!=nil{
		log.Fatal(err)
	}
	
	if searchIndex(mapCreated,word){
		fmt.Printf("OLD Map : %v",mapCreated)
		delete(mapCreated,word)
		fmt.Printf("\nNew Map : %v",mapCreated)
	}else{
		fmt.Printf("Not Founded ! Sorry !")
	}
}