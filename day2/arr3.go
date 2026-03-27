package main
import (
	"fmt"

)

func squares(n int)[]int{
result:= make([]int, n)
	for i:=0; i<n; i++{
		result[i]= (i+1)*(i+1)
	}
	return result

}
func main() {

	result:= squares(5)
	fmt.Println(result)
}
