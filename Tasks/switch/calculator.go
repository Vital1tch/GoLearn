package _switch

import "fmt"

func Calculator(number1 float64, number2 float64, mark string) {
	switch mark {
	case "+":
		result := number1 + number2
		fmt.Printf("%.2f + %.2f = %.2f\n", number1, number2, result)

	case "-":
		result := number1 - number2
		fmt.Printf("%.2f - %.2f = %.2f\n", number1, number2, result)

	case "*":
		result := number1 * number2
		fmt.Printf("%.2f * %.2f = %.2f\n", number1, number2, result)

	case "/":
		result := number1 / number2
		if number2 == 0 {
			fmt.Println("Делить на ноль нельзя!")
			return
		}
		fmt.Printf("%.2f / %.2f = %.2f\n", number1, number2, result)

	default:
		fmt.Println("Убедитесь, что вы ввели допустимые знаки действий!")
	}
}
