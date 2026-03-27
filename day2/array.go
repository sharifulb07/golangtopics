package main 
import (
	"fmt"
)

func ArrExample()[2][3]int{
	var matrix[2][3]int=[2][3]int{{1,2,3},{4,5,6}}
	return matrix
}

func main(){
	matrix := ArrExample()

	for i:= 0; i < len(matrix); i++ {
		for j:=0; j < len(matrix[i]); j++ {
			fmt.Printf("%d ", matrix[i][j])
		}
}
}
