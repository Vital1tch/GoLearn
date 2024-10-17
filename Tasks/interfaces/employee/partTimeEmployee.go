package employee

type PartTimeEmployee struct {
	Hours        float64
	SalaryOfHour float64
}

func (e *PartTimeEmployee) GetSalary() float64 {
	return e.Hours * e.SalaryOfHour
}
