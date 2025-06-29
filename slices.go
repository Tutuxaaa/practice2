package main

import "fmt"

// Поиск
func Find(slice []string, target string) (int, bool) {
	for i, v := range slice {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

// Быстрая сортировка
func QuickSort(slice []string) []string {
	if len(slice) < 2 {
		return slice
	}

	left, right := 0, len(slice)-1
	pivot := len(slice) / 2

	slice[pivot], slice[right] = slice[right], slice[pivot]

	for i := range slice {
		if slice[i] < slice[right] {
			slice[left], slice[i] = slice[i], slice[left]
			left++
		}
	}

	slice[left], slice[right] = slice[right], slice[left]

	QuickSort(slice[:left])
	QuickSort(slice[left+1:])

	return slice
}

// Фильтрация (элементы, у которых четных индекс)
func Filter(slice []string) []string {
	var result []string
	for i, v := range slice {
		if i%2 == 0 {
			result = append(result, v)
		}
	}
	return result
}

func main() {
	data := []string{"hello", "my", "name", "is", "Veronika"}

	if idx, found := Find(data, "my"); found {
		fmt.Printf("Найдено по индексу %d\n", idx)
	}

	// Сортировка
	sorted := QuickSort(data)
	fmt.Println("Отсортировано:", sorted)

	// Фильтрация
	evenNumbers := Filter(data)
	fmt.Println("Элементы среза, которые стоят на четных позициях", evenNumbers)

}
