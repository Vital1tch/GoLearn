package panics

import "fmt"

func divide(number1, number2 int) int {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered by panic", r)
		}
	}()

	if number2 == 0 {
		panic("Division by zero!")
	}
	return number1 / number2
}
