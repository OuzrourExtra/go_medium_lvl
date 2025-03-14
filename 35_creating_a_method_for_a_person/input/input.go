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
