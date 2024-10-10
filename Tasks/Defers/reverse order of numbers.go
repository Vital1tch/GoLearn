package Defers

import "fmt"

func reverseNumbers() {
	i := 1
	for ; i <= 5; i++ {
		fmt.Println(i)
		defer fmt.Println(i)
	}
}
