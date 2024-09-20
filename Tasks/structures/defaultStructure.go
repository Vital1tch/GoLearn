package structures

import "fmt"

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{name, age}
}

func Greet(p *Person) {
	fmt.Printf("Привет! Меня зовут %s, мне %d \n", p.Name, p.Age)
}
