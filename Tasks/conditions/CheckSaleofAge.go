package conditions

import "fmt"

func CheckSaleOfAge(summ float64, age int) {
	//Проверка на сумму более 10000
	if summ >= 10000 {
		summ -= summ * 0.1
	}
	//Проверка на возраст если меньше 18 или возраст больше 60
	if age <= 18 || age >= 60 {
		summ -= summ * 0.05
	}
	fmt.Printf("Итоговая сумма вашей покупки: %.2f", summ)
}
