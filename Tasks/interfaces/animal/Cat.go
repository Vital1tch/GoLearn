package animal

type Cat struct {
	Name  string
	Voice string
}

func (c Cat) Speak() string {
	return c.Voice
}

func (c Cat) GetName() string {
	return c.Name
}
