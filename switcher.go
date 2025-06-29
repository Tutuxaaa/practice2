package main

import "fmt"

// Функция, принимающая параметр по значению (копия)
func modifyValue(val int) {
	val = val + 10
	fmt.Println("Inside modifyValue:", val)
}

// Функция, принимающая параметр по указателю (ссылка)
func modifyPointer(ptr *int) {
	*ptr = *ptr + 10
	fmt.Println("Inside modifyPointer:", *ptr)
}

func main() {
	original := 5

	fmt.Println("Before modifyValue:", original)
	modifyValue(original)
	fmt.Println("After modifyValue:", original)

	fmt.Println("--------")

	fmt.Println("Before modifyPointer:", original)
	modifyPointer(&original)
	fmt.Println("After modifyPointer:", original)
}
