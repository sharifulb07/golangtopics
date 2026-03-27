
package main 
import (
	"fmt"
)


type MyError struct{
	Code int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}


func Test(e MyError){
	fmt.Println(e.Error())
}




func main(){
	Test(MyError{Code: 400, Message: "Not Found"})
}