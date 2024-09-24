package _switch

import "fmt"

func DisplayGrade(grade int) {
	switch {
	case grade >= 90 && grade <= 100:

		fmt.Println("Ваша оценка отлично!")

	case grade >= 75 && grade <= 89:
		fmt.Println("Ваша оценка хорошо!")

	case grade >= 60 && grade <= 74:
		fmt.Println("Ваша оценка удовлетворительно!")

	case grade >= 0 && grade <= 59:
		fmt.Println("Ваша оценка неудовлетворительно!")

	default:
		fmt.Println("Введите число от 100 до 0")
	}
}
