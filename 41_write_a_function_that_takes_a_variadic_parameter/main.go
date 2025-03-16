package main

import (
	"fmt"
	"bufio"
	"strings"
	"os"
	"strconv"
	"log"
)


/*
inputNumber is a function for inputing a number
*/
func InputNumber(sentence string) (nb float64) {
	fmt.Println("=======================")
	fmt.Printf("%s: ",sentence)
	reader := bufio.NewReader(os.Stdin)
	str, errReader := reader.ReadString('\n')
	if errReader != nil{
		log.Fatal(errReader)
	}
	str = strings.TrimSpace(str)
	nb, errConv := strconv.ParseFloat(str, 64)
	if errConv != nil {
		log.Fatal(errConv)
	}
	return nb
}


func average(notes ...float64) float64{
	n , somme := 0.0 , 0.0
	for _ , note := range notes{
		somme += note
		n++
	}
	return somme / n
}

func main(){
	numberOfNotes := int(InputNumber("enter the number of notes"))
	Notes := make([]float64,int(numberOfNotes))
	for i:=0;i<numberOfNotes;i++{
		Notes[i] = InputNumber("Enter the note")
	}
	averageOfNotes := average(Notes...)
	fmt.Printf("===========================================")
	fmt.Printf("\nThe Average of the class is %0.2f",averageOfNotes)
	fmt.Printf("\n===========================================")
	
}