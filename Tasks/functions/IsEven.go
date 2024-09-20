package functions

func IsEven(number int) string {
	if number%2 == 0 {
		return "Четное"
	}
	return "Нечётное"
}
