package main

import (
	"errors"
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
func InputNumber(sentence string) (nb int64, err error) {
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



func multipleReturns(number int) (result int , err error){
	if number < 10 {
		return number , nil
	}else{
		return 0 , errors.New("the Number isn't Greater than 10")
	}
}


func main(){
	number , err := InputNumber("Enter the number you want")
	if err!=nil{
		log.Fatal(err)
	}
	numberCheck , err2 := multipleReturns(int(number))
	if err2 != nil{
		log.Fatal(err2)
	}
	fmt.Printf("========\n NICE : %d < 10 ",numberCheck)

}
