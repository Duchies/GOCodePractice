package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var i = 10

	strConv := strconv.Itoa(i)

	// integer is converted to string
	fmt.Println(strConv)

	var ch []rune
	ch = append(ch, 'a', 'b', 'c', 'd', 'e', 'f', 'g')

	result := string(ch)
	fmt.Println(reflect.TypeOf(result))
	fmt.Println(result)

	ch[0] = 'z'
	
	
	// achive mutability in string with rune 
	result = string(ch)
	fmt.Println(result)

}
