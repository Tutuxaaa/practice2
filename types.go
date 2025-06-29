package main

import (
	"fmt"
	"unsafe"
)

var typeInt = 128
var typeFloat = 2.7
var typeComplex complex64 = 1 + 2i
var typeByte = 1
var typeString = "Hello"
var typeBool = true

func main() {
	fmt.Println(unsafe.Sizeof(typeInt))
	fmt.Println(unsafe.Sizeof(typeFloat))
	fmt.Println(unsafe.Sizeof(typeComplex))
	fmt.Println(unsafe.Sizeof(typeByte))
	fmt.Println(unsafe.Sizeof(typeString))
	fmt.Println(unsafe.Sizeof(typeBool))
}
