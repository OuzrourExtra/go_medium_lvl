package input

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


/*
inputString is a function for inputing a string
*/
func InputString(sentence string) (strResult string) {
	fmt.Println("=======================")
	fmt.Printf("%s: ",sentence)
	reader := bufio.NewReader(os.Stdin)
	str, errReading := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	if errReading != nil {
		log.Fatal(errReading)
		return ""
	} else {
		return str
	}
}
