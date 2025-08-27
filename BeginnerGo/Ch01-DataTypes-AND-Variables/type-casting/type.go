package main

import (
	"fmt"
	"reflect"
)

// fmt : Formart + Print
// reflect : to get the type of a variable
// strconv : to convert string to int and vice versa: string convert

func main() {
	msg := "Hello World"
	pi := 3.14

	fmt.Println(reflect.TypeOf(msg))
	fmt.Println(reflect.TypeOf(pi))

}
