package main

import (
	"fmt"
	"log"
	"strings"
	"bufio"
	"os"
)

func detectRepetition(String string) (map[string]int){
	words := strings.Fields(String)
	mapOfOccurence := make(map[string]int)
	for _ , value := range words{
		mapOfOccurence[value]++
	}
	return mapOfOccurence
}


func main(){
	fmt.Printf("Enter the Sentence : ")
	reader := bufio.NewReader(os.Stdin)
	sentence , err := reader.ReadString('\n')
	if err!=nil{
		log.Fatal(err)
	}
	mapOfOccurence := detectRepetition(sentence)
	fmt.Println("\n=================================")
	fmt.Println("List Of Repetition")
	fmt.Println("=================================")
	for str , numberOfRepetition := range mapOfOccurence{
		fmt.Printf("word '%s' : %d time\n",str,numberOfRepetition)
	}
	
}