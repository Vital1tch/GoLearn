package slices

func EvenSlice(arr []int) []int {
	arr2 := []int{} //Пустой срез, в который мы будем записывать данные

	for index, value := range arr {
		arr[index] = value //Лишняя строка, но это присвоение элемента
		if value%2 == 0 {  //Сравниваем, чётное ли число
			arr2 = append(arr2, value) // Записаваем в срез получившиеся значения
		}
	}
	return arr2
}
