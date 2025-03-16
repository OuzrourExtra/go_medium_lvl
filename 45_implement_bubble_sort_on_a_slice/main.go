package main

import (
	"fmt"
	"math/rand"
	"time"
)


func BubbleSort(slice *[]int){
	for i:=0 ; i<len(*slice) ; i++{
		swapped := false
		for j:=0 ; j<len(*slice)-i-1 ; j++{
			if (*slice)[j] > (*slice)[j+1] {
				(*slice)[j] , (*slice)[j+1] = (*slice)[j+1] , (*slice)[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func main(){
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	slice := make([]int , 50)
	for i:=0 ; i<50 ; i++{
		slice[i]=r.Intn(10000)
	}
	fmt.Printf("\nSlice Before Sorting : %v",slice)
	fmt.Printf("\n=======================")
	BubbleSort(&slice)
	fmt.Printf("\nSlice After Sorting : %v",slice)
	fmt.Printf("\n=======================")


}

