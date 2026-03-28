package main 
import (
	"fmt"
)


func main(){

	slice:=make([]int,5,5)

	dist:=make([]int, len(slice))

	copy(dist, slice)

	dist[0]=100
	dist[1]=200
	dist[2]=300
	
	fmt.Println("Original slice:", slice)
	fmt.Println("Copied slice:", dist)

}