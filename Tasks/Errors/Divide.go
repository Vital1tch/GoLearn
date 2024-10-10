package Errors

import (
	"errors"
	"fmt"
)

func divide(number1, number2 int) (int, error) {
	if number2 == 0 {
		return 0, errors.New("Cannot divide by zero")
	} else {
		return number1 / number2, nil
	}
}
func main() {
	number1, number2 := 10, 2
	result, err := divide(number1, number2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
