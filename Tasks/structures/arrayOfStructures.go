package structures

import "fmt"

type Product struct {
	NameofProduct string
	Price         float64
}
type ArrayProduct struct {
	ProductList []Product
}

func ShowArrayProduct(p *ArrayProduct) {
	for _, product := range p.ProductList {
		fmt.Printf("Продукт %s, Цена: %.2f\n", product.NameofProduct, product.Price)
	}
}
