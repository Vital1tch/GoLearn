package methods

import "fmt"

type Counter struct {
	value int
}

func CreateNewCounter(value int) *Counter {
	return &Counter{
		value,
	}
}

func (c *Counter) Increment() {
	c.value += 1
}

func (c *Counter) ShowValue() {
	fmt.Printf("Counter ShowValue : %d\n", c.value)
}
