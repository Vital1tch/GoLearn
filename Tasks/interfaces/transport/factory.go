package transport

import (
	"GoLearn/Tasks/interfaces/transport/types"
	"fmt"
)

// Фабрика для создания экземпляров Ship или Truck
func transportFactory(transportType string, params ...interface{}) Transport {
	switch transportType {
	case "Ship":
		// Убеждаемся, что переданы правильные параметры
		if len(params) != 3 {
			fmt.Println("Ошибка: недостаточно параметров для создания Ship")
			return nil
		}
		if capacity, ok1 := params[0].(int64); ok1 {
			if currentLocation, ok2 := params[1].(string); ok2 {
				if destination, ok3 := params[2].(string); ok3 {
					return types.NewShip(capacity, currentLocation, destination)
				}
			}
		}

	case "Truck":
		if len(params) != 3 {
			fmt.Println("Ошибка: недостаточно параметров для создания Truck")
			return nil
		}
		if load, ok1 := params[0].(int64); ok1 {
			if fuel, ok2 := params[1].(float64); ok2 {
				if destination, ok3 := params[2].(string); ok3 {
					return types.NewTruck(load, fuel, destination)
				}
			}
		}

	}
	fmt.Printf("Ошибка: неверные параметры для создания %s\n", transportType)
	return nil
}
