package functions

func Apply(fn func(int) int, value int) int {
	return fn(value)
}

func Calc(fn func(int, int) int) int {
	return fn(10, 10)
}
