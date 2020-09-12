// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type User struct {
// 	Name    string
// 	Address string
// }

// func main() {
// 	fmt.Println("Please type your input floating number, thanks!")
// 	f := map[string]string{}
// 	var inputName string
// 	var inputAddress string

// 	fmt.Printf("Insert your name:\n")
// 	fmt.Scanf("%s\n", &inputName)
// 	fmt.Printf("Insert your address:\n")
// 	fmt.Scanf("%s", &inputAddress)

// 	// the first one
// 	f["name"] = inputName
// 	f["address"] = inputAddress
// 	byteResult, err := json.Marshal(f)

// 	// the second one
// 	// user := User{
// 	// 	Name:    inputName,
// 	// 	Address: inputAddress,
// 	// }
// 	// byteResult, err := json.Marshal(user)

// 	if err != nil {
// 		panic("Can't convert object to json")
// 	}

// 	fmt.Println(string(byteResult))
// }
