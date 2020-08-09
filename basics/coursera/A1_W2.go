package main

import (
	"fmt"
	_ "fmt"
)

func main() {
	fmt.Println("Please type your input floating number, thanks!")

	var inputFloat float64
	fmt.Scanf("%f", &inputFloat)
	i := int(inputFloat)

	fmt.Printf("The integer of the user input floating number %f is %d\n", inputFloat, i)
}
