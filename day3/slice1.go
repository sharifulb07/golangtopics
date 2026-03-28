// Online Go compiler to run Golang program online
// Print "Try programiz.pro" message

package main

import (
	"fmt"

)

func main() {
 sl:=make([]int, 0)
 
 for i:=0;i<10;i++{
     sl=append(sl,i)
 }
 
 fmt.Printf("length: %d, cap:%d, slice:%v \n", len(sl), cap(sl), sl)
}