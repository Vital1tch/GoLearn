package conditions

import "fmt"

type Person struct {
	Name string
	Age  int
}

func CreatePerson(name string, age int) *Person {
	return &Person{
		Name: name,
		Age:  age,
	}
}

func (p *Person) CheckForParityAge() {
	if p.Age%2 == 0 {
		fmt.Printf("Возраст %s чётный\n", p.Name)
	} else {
		fmt.Printf("Возраст %s нечётный\n", p.Name)
	}
}
