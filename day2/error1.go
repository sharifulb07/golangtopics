package main 
import (
	"fmt"
)


func process(data string) error{

	if data ==""{
		fmt.Println("Data cannot be Empty")
	}

	if len(data)< 5{
		fmt.Println("Length of data is too short")
	}

	// main logic

	return nil

}

func main()  {
	err := process("Shariful")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Data processed successfully")
	}
}