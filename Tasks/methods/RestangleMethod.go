package methods

import "fmt"

type Restangle struct {
	width, height float64
}

func (r *Restangle) Area() float64 {
	return r.height * r.width
}

func NewRestangle(width, height float64) *Restangle {
	return &Restangle{width, height}
}

func ShowRestangle(r *Restangle) { //Функция
	fmt.Printf("Ширина: %f\n Длина: %f\n", r.width, r.height)
}

func (r *Restangle) ShowRes() { //Метод только для структуры Restangle
	fmt.Printf("Ширина: %f\n Длина: %f\n", r.width, r.height)
}
