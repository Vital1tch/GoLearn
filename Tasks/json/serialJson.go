package json

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ID       int     `json:"ID"`
	Name     string  `json:"Name"`
	Price    float64 `json:"Price"`
	InStock  bool    `json:"is_available"`
	Category string  `json:"Category,omitempty"`
}

func serialJson(products []Product) []byte {
	productData, err := json.MarshalIndent(products, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(productData))

	return productData
}
