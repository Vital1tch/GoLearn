package json

import (
	"encoding/json"
	"fmt"
)

func deSerialJson(productData []byte) {
	var newDataOfProducts []Product
	err := json.Unmarshal(productData, &newDataOfProducts)
	if err != nil {
		panic(err)
	}

	for index := range newDataOfProducts {
		if newDataOfProducts[index].Category == "" {
			newDataOfProducts[index].Category = "N/A"
		}
	}

	fmt.Printf("Десериализованный json: %+v\n", newDataOfProducts)
}
