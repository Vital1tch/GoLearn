package loops

func Fibonachi(n int) int {
	if n < 2 {
		return n
	}
	return Fibonachi(n-1) + Fibonachi(n-2)
}
