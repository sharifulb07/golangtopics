package main 

import (
	"fmt"
	"encoding/json"
	"log"
)


type User struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Email string `json:"email, omitempty"`
	IsAdmin bool `json:"is_admin"`

}




func main(){
	u1:=User{
		ID: 1,
		Name: "Shariful",
		Email: "example@gmail.com",
		IsAdmin: true,
	}


	// json. marshal
	data, err:=json.Marshal(u1)

	if err !=nil{
		log.Fatal(err)

	}




	fmt.Println("Data is String format: ", string(data))
	// fmt.Println(data[len(data)-1])
	// fmt.Println(append(data, 45))
	// fmt.Println(len(string(data)))



	// json unmarshal

jsonStr := `{
        "id": 100,
        "full_name": "Bob Ahmed",
        "is_admin": false
    }`

		var u2 User


		// if err:=json.Unmarshal([]byte(jsonStr),&u2); err!=nil{
		// 	log.Fatal(err)
		// }

		// fmt.Printf("Unmarshal: %+v\n", u2)

		// fmt.Printf("Unmarshaled: %+v \n", u2)

		if err := json.Unmarshal([]byte(jsonStr), &u2); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Unmarshaled: %+v\n", u2)
}