package main

import "fmt"

//находит самую длинную строку из переданных
func LongestString(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]

	for _, s := range strings[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}

	return longest
}

func main() {
	// Примеры использования
	fmt.Println(LongestString("cat", "dog", "elephant"))
	fmt.Println(LongestString("apple", "banana", "pear"))
	fmt.Println(LongestString("one", "two", "three", "four"))
	fmt.Println(LongestString("longer", "longest"))
	fmt.Println(LongestString(""))
}
