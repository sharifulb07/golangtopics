package main 
import (
	"fmt"
)


func main(){

	arr:=[6]int {2,3,4,5,6,7}

	slice1:=arr[1:4]
	slice2:=arr[:3]
	slice3:=arr[3:]

	fmt.Println("Original array:", arr)
	fmt.Println("Slice 1 (arr[1:4]):", slice1)
	fmt.Println("Slice 2 (arr[:3]):", slice2)
	fmt.Println("Slice 3 (arr[3:]):", slice3)

	slice3[0]=10
	fmt.Println("After modifying slice2:", slice2)
	fmt.Println("Original array after modifying slice2:", arr)
}