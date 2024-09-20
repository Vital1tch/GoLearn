package Pointers

type Person struct {
	Name string
	Age  int
}

func UpdatePerson(p *Person, newName string, newAge int) {
	p.Name = newName
	p.Age = newAge
}
