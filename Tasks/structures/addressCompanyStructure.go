package structures

import "fmt"

type Address struct {
	Town   string
	Street string
}

type Company struct {
	NameOfCompany string
	Address       //Вложенная структура
}

func NewCompany(nameOfCompany, town, street string) *Company {
	return &Company{
		NameOfCompany: nameOfCompany,
		Address: Address{
			Town:   town,
			Street: street,
		},
	}
}

func ShowInformationOfCompany(c *Company) {
	fmt.Printf("Company: %s\n Town: %s\n Street:%s\n ", c.NameOfCompany, c.Address.Town, c.Address.Street)
}
