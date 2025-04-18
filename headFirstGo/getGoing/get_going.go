package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello Go")
	fmt.Println(reflect.TypeOf(32))
	fmt.Println(reflect.TypeOf("ABC"))
	fmt.Println(reflect.TypeOf(43.67))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf('A'))
}
