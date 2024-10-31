package transport

import "fmt"

func DeliverAll(transports []Transport) {
	for _, transport := range transports {
		if transport == nil {
			fmt.Println("Ошибка: обнаружен nil в списке транспортных средств.")
			continue
		}
		transport.Deliver()
		fmt.Printf("Скорость доставки: %d км/ч \n", transport.Speed())
	}
}
