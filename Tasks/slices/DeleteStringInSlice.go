package slices

import "fmt"

func DeleteStringInSlice(slice []string, s string) ([]string, []string) {
	for index, value := range slice {
		if value == s { //Поиск строки-аргумента в массиве
			copy(slice[index:], slice[index+1:])
			slice[len(slice)-1] = ""
			slice = slice[:len(slice)-1]
		}
	}
	fmt.Printf("Удалённая строка:%s\n", s)
	return slice, nil
}
