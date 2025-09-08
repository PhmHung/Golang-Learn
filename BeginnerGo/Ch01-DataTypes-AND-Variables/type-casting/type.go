package main

import (
	"fmt"
	"reflect"
	"strconv"
)

// fmt : Formart + Print
// reflect : to get the type of a variable
// strconv : to convert string to int and vice versa: string convert

func main() {
	msg := "Hello World"
	pi := 3.14

	fmt.Println(reflect.TypeOf(msg))
	fmt.Println(reflect.TypeOf(pi))

	//convert int, float64, float32
	var percent int = 70
	fmt.Println(float64(percent))
	fmt.Println(int(percent))

	//int to string
	var s string = strconv.Itoa(percent)
	fmt.Println(reflect.TypeOf(s))

	//string to int
	// Fail to parsing a real string to an int
	out, error1 := strconv.Atoi(msg)
	fmt.Println(out)
	fmt.Println(error1)

	//Parsing success
	newNum, error2 := strconv.Atoi(s)
	fmt.Println(newNum)
	fmt.Println(error2)
}
