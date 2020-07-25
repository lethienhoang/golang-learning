package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name   string `json:"name`
	Type   string `json:"type`
	Age    int    `json:Age`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	// read file into memory
	jsonFile, err := os.Open("users.json")

	// check error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully opened user.json")

	// defer statements delay the execution of the function or method or an anonymous method until the nearby functions returns
	defer jsonFile.Close()
	// convert jsonFile to by array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}
