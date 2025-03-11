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

// commonElementsDetector is a function that detect the common
// Elements between 2 slices
func commonElementsDetector(slice1 []int , slice2 []int) ([]int){
	common := make(map[int]bool)
	commonSlice := make([]int,0,len(slice1))

	for _ , value := range(slice1){
		common[value] = true
	}
	for _ , value := range(slice2){
		if common[value]{
			commonSlice = append(commonSlice, value)
			// assign false to the value to avoid duplication
			common[value]=false
		}
	}
	return commonSlice
}

func appendSlice(slice *[]int , capacity int){
	for i:= 0 ; i<capacity ; i++{
			new , err:= inputNumber("Enter an elements of Slice")
			if err!= nil{
				log.Fatal(err)
			}
		*slice = append(*slice , int(new))
	}
}

func main(){
	// Slice 1 

	numberSlice1, err1 := inputNumber("Enter the number of elements in the slice 1")
	if err1!=nil{
		log.Fatal(err1)
	}
	slice1 := make([]int , 0 ,  numberSlice1)
	appendSlice(&slice1,int(numberSlice1))

	// Slice 2

	numberSlice2 , err2 := inputNumber("Enter the number of elements in the slice 2")
	if err2!=nil{
		log.Fatal(err2)
	}
	slice2 := make([]int, 0 , numberSlice2)
	appendSlice(&slice2,int(numberSlice2))

	// Merged Slice : 

	commonElementsSlice := commonElementsDetector(slice1,slice2)
	fmt.Printf("\n ===========================")
	fmt.Printf("\n The Merged Slice : %v",commonElementsSlice)

}