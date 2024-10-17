package animal

type Cow struct {
	Name  string
	Voice string
}

func (c Cow) Speak() string {
	return c.Voice
}
func (c Cow) GetName() string {
	return c.Name
}
