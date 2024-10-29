package buffer

import (
	"bytes"
	"fmt"
)

type Product struct {
	Name  string
	price float64
}

func ProductList(products []Product) {
	var buf bytes.Buffer

	for _, product := range products {
		buf.WriteString(product.Name)
		buf.WriteString(": ")
		price := fmt.Sprintf("$%.2f\n", product.price)
		buf.WriteString(price)
	}
	fmt.Println(buf.String())
}
