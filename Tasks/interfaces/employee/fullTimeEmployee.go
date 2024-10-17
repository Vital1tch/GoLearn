package employee

type FullTimeEmployee struct {
	Salary float64
}

func (e *FullTimeEmployee) GetSalary() float64 {
	return e.Salary
}
