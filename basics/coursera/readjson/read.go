package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	maxLength = 20
)

type Users struct {
	Users []Infor `json:"users"`
}

type Infor struct {
	Fname string `json:"first_name"`
	Lname string `json:"last_name"`
}

func (u *Infor) Set() {
	if len(u.Fname) > maxLength {
		u.Fname = u.Fname[:maxLength]
	}

	if len(u.Lname) > maxLength {
		u.Lname = u.Lname[:maxLength]
	}
}

func main() {
	var fileName string
	fmt.Printf("Enter file name:\n")
	fmt.Scanf("%s", &fileName)

	fileName = fileName + ".json"
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully opened %s.json \n", fileName)

	defer file.Close()

	byteValue, _ := ioutil.ReadAll(file)

	var users Users

	json.Unmarshal(byteValue, &users)

	for i := 0; i < len(users.Users); i++ {
		users.Users[i].Set()

		fmt.Println("First name: " + users.Users[i].Fname)
		fmt.Println("Last name: " + users.Users[i].Lname)
		fmt.Printf("********* \n")
	}

}
