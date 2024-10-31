package transport

type Transport interface {
	Deliver()
	Speed() int
}
