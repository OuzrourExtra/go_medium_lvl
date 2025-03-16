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
func InputNumber(sentence string) (nb int) {
	fmt.Println("=======================")
	fmt.Printf("%s: ",sentence)
	reader := bufio.NewReader(os.Stdin)
	str, errReader := reader.ReadString('\n')
	if errReader != nil{
		log.Fatal(errReader)
	}
	str = strings.TrimSpace(str)
	nb, errConv := strconv.Atoi(str)
	if errConv != nil {
		log.Fatal(errConv)
	}
	return nb
}

func fib(n int)int {
	if n==0 || n==1 {
		return 1
	}else{
		return fib(n-1)+fib(n-2)
	}
}

func main(){
	number := InputNumber("Enter the number of fibonnaci")
	for i:=0 ; i<=number ; i++{
		fmt.Printf("\n Fibonnaci(%d)=%d ",i,fib(i))
	}
}