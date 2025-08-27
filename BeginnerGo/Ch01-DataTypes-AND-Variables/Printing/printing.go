package main

import "fmt"

func main() {
	fmt.Print("Hello", "\n")
	fmt.Println("The film", 12, "chapter", true)

	//Printf :special format
	// %v -> default format
	// %T -> type of the value
	// %d -> integers
	// %c -> character
	// %q -> quoted character / string
	// %s -> plain string
	// %t -> true or false
	// %f -> floating numbers
	// %0.2f -> floating numbers up to 2 decimal places

	var value float32 = 3.141256
	var name string = "PI Number"
	var boolean bool = true
	var radius int = 3
	fmt.Printf("The value of %q is %f and 2 decimal places of %q is %0.2f, the radius of the circle if %d, the answer of this question is %t", name, value, name, value, radius, boolean)
	fmt.Print("\n")
}
