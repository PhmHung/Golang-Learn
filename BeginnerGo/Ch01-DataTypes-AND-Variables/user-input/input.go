package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var firstName string
	var fullName string
	var age int
	var class string

	//fmt.Scanln(&str) get only one string
	fmt.Println("Enter your first name: ")
	fmt.Scanln(&firstName)
	fmt.Println("Your first name is: ", firstName)

	//Using bufio same as getln(cin, str) for get all string with space
	fmt.Println("Enter your full name: ")
	reader := bufio.NewReader(os.Stdin)
	fullName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ", err)
	}
	fullName = strings.TrimSpace(fullName)
	fmt.Println("Your full name is: ", fullName)

	// Get two string at the same time with fmt.Scanln(&str1,&str2)
	fmt.Println("Enter your age , and class name in field: ")
	fmt.Scanln(&age, &class)
	fmt.Println("Your age: ", age)
	fmt.Println("Your class name: ", class)

	// Enter your first name:
	// Hung
	// Your first name is:  Hung
	// Enter your full name:
	// Hung Pham
	// Your full name is:  Hung Pham
	// Enter your age , and class name in field:
	// 21 IT20
	// Your age:  21
	// Your class name:  IT20
}
