package main

import "fmt"

// global variables
const pi float64 = 3.14
const sum = 3

var a string = "Dog"

func main() {

	//Var in block, global var, local var
	var b string = "bar"
	{
		var c string = "full-snack"
		fmt.Println(b)
		fmt.Println(c)
		fmt.Println(a)
	}

	fmt.Println(b)
	fmt.Println(pi)
	fmt.Printf("Type of sum is: %T", sum)
}
