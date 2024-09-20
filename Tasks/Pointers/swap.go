package Pointers

func Swap(a, b *int) {
	*a, *b = *b, *a
}
