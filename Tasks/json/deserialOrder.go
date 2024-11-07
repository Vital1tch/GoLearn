package json

import "encoding/json"

type User struct {
	ID    int    `json:"ID"`
	Name  string `json:"Name"`
	Age   int    `json:"Age"`
	Email string `json:"Email"`
}

type Product2 struct {
	Name  string  `json:"Name"`
	Price float64 `json:"Price"`
}

type Order struct {
	OrderID  int        `json:"OrderID"`
	Customer User       `json:"Customer"`
	Items    []Product2 `json:"Items"`
}

func deSerialOrder(orderData []byte) (Order, error) {
	var newOrder Order
	err := json.Unmarshal(orderData, &newOrder)
	if err != nil {
		return Order{}, err
	}
	return newOrder, nil
}
