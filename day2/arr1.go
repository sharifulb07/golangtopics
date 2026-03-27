package main 
import (
	"fmt"
)


func modify(arr *[3]int){

	for i:=range arr{
		arr[i]=arr[i]*3
	}
}


func main(){
	arr:=[3]int{1,2,3}
	fmt.Println("Before modification:", arr)
modify(&arr)
fmt.Println("After modification:", arr)
}