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
inputString is a function for inputing a string
*/
func InputString(sentence string) (strResult string, err error) {
	fmt.Println("=======================")
	fmt.Printf("%s: ",sentence)
	reader := bufio.NewReader(os.Stdin)
	str, err1 := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	if err1 != nil {
		log.Fatal(err1)
		return "", err1
	} else {
		return str, nil
	}
}


func StringToInt(str string) int{
	number , err := strconv.Atoi(str)
	if err != nil {
		log.Fatal(err)
	}
	return number
}

func main(){
	str , err := InputString("Enter the Sentence")
	if err != nil{
		log.Fatal(err)
	}
	number := StringToInt(str)
	fmt.Printf("===============\n the Number : %d !",number)
}