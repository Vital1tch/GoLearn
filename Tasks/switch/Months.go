package _switch

import "fmt"

func FindMonth(numberOfMonth int) {
	switch numberOfMonth {
	case 1:
		fmt.Println("Январь - 31 день")

	case 2:
		fmt.Println("Февраль - 28/29 дней")

	case 3:
		fmt.Println("Март - 31 день")

	case 4:
		fmt.Println("Апрель - 30 дней")

	case 5:
		fmt.Println("Май - 31 день")

	case 6:
		fmt.Println("Июнь - 30 дней")

	case 7:
		fmt.Println("Июль - 31 день")

	case 8:
		fmt.Println("Август - 31 день")

	case 9:
		fmt.Println("Сентябрь - 30 дней")

	case 10:
		fmt.Println("Октябрь - 31 день")

	case 11:
		fmt.Println("Ноябрь - 30 дней")

	case 12:
		fmt.Println("Декабрь - 31 день")

	default:
		fmt.Println("Ошибка! Убедитесь, что ввели числа в диапазоне от 1 - 12")
	}

}
