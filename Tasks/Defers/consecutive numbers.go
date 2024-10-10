package Defers

import "fmt"

func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
func third() {
	fmt.Println("Third")
}

//func main() {
//	defer first()
//	defer second()
//	defer third()
//}
