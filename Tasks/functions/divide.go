package functions

import "errors"

func Divide(number1, number2 float64) (float64, error) {
	if number2 == 0 {
		return 0, errors.New("Деление на ноль не возможно ")
	}
	result := number1 / number2
	return result, nil
}
