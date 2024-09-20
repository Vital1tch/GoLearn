package functions

func Counter() func() int {
	number := 0
	return func() int {
		number++
		return number
	}
}
