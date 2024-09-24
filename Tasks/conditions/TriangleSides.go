package conditions

import "fmt"

func CheckTriangleSides(firstSide, secondSide, thirdSide float64) {
	//Проверка на существование треугольника
	if firstSide+secondSide > thirdSide && thirdSide+firstSide > secondSide && secondSide+thirdSide > firstSide {

		//Проверка на равностороний
		if firstSide == thirdSide && secondSide == thirdSide {
			fmt.Println("Треугольник равносторонний")
			return
		} else if firstSide == thirdSide || secondSide == thirdSide || firstSide == secondSide {
			fmt.Println("Треугольник равнобедренный")
			return
		} else {
			fmt.Println("Треугольник разносторонний")
		}

	} else {
		fmt.Println("Ошибка! Треугольника не существует")
	}
}
