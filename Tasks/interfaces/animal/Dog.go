package animal

type Dog struct {
	Name  string
	Voice string
}

func (d Dog) Speak() string {
	return d.Voice
}

func (d Dog) GetName() string {
	return d.Name
}
