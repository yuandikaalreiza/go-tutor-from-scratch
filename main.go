package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonString = `{"Name": "john wick", "Age": 27}`
	var jsonData = []byte(jsonString)

	var data User

	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age :", data.age)
}

type User struct {
	FullName string `json:"name"`
	age      int
}
