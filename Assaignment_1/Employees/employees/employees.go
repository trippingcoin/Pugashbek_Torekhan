package employees

import "fmt"

type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

func (fte FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("FullTimeEmployee - ID: %d, Name: %s, Salary: %d Tenge", fte.ID, fte.Name, fte.Salary)
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (pte PartTimeEmployee) GetDetails() string {
	totalPay := float64(pte.HourlyRate) * float64(pte.HoursWorked)
	return fmt.Sprintf("PartTimeEmployee - ID: %d, Name: %s, HourlyRate: %d Tenge, HoursWorked: %.2f, Total Pay: %.2f Tenge",
		pte.ID, pte.Name, pte.HourlyRate, pte.HoursWorked, totalPay)
}

type Company struct {
	Employees map[string]Employee
}

func NewCompany() *Company {
	return &Company{Employees: make(map[string]Employee)}
}

func (c *Company) AddEmployee(emp Employee) {
	var id string
	switch e := emp.(type) {
	case FullTimeEmployee:
		id = fmt.Sprintf("%d", e.ID)
	case PartTimeEmployee:
		id = fmt.Sprintf("%d", e.ID)
	}
	c.Employees[id] = emp
	fmt.Println("Employee added successfully!")
}

func (c *Company) ListEmployees() {
	if len(c.Employees) == 0 {
		fmt.Println("No employees in the company.")
		return
	}

	fmt.Println("Employee List:")
	for _, emp := range c.Employees {
		fmt.Println(emp.GetDetails())
	}
}
