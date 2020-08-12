package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func foundCharacterInString(s string) string {
	text := strings.Split(s, " ")

	arr := []string{}
	for i := 0; i < len(text); i++ {
		lenOfText := len(text[i])
		lenOfText = lenOfText - 1

		if strings.Index(text[i], "i") == 0 && strings.LastIndex(text[i], "n") == lenOfText && strings.Contains(text[i], "a") {
			arr = append(arr, "Found")
		} else {
			arr = append(arr, "Not")
		}
	}

	for i := 0; i < len(arr); i++ {
		matched, _ := regexp.MatchString(arr[i], "Not")

		if matched {
			return "Not Found!"
		}
	}

	return "Found!"
}

func main() {
	fmt.Println("Welcome")

	var inputString string

	fmt.Println("Please input text")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	inputString = scanner.Text()

	inputString = strings.ToLower(inputString)

	fmt.Println("Your input is convert to lowser case ", inputString)

	result := foundCharacterInString(inputString)

	fmt.Println("The result ", result)
}
