package main

import (
	"fmt"
	"strings"

	"github.com/trippingcoin/Pugashbek_Torekhan/Assaignment_1/Assaignment_1/Employees/employees"
)

func main() {
	company := employees.NewCompany()

	for {
		fmt.Println("Choose an option: Add (Employee), List (Employees), Exit")
		var command string
		fmt.Print("> ")
		fmt.Scanln(&command)

		switch strings.ToLower(command) {
		case "add":
			fmt.Println("Select Employee Type: FullTime, PartTime")
			var empType string
			fmt.Print("> ")
			fmt.Scanln(&empType)

			switch strings.ToLower(empType) {
			case "fulltime":
				var id uint64
				var name string
				var salary uint32

				fmt.Print("Enter ID: ")
				fmt.Scanln(&id)
				fmt.Print("Enter Name: ")
				fmt.Scanln(&name)
				fmt.Print("Enter Salary: ")
				fmt.Scanln(&salary)

				emp := employees.FullTimeEmployee{
					ID:     id,
					Name:   name,
					Salary: salary,
				}
				company.AddEmployee(emp)

			case "parttime":
				var id uint64
				var name string
				var hourlyRate uint64
				var hoursWorked float32

				fmt.Print("Enter ID: ")
				fmt.Scanln(&id)
				fmt.Print("Enter Name: ")
				fmt.Scanln(&name)
				fmt.Print("Enter Hourly Rate: ")
				fmt.Scanln(&hourlyRate)
				fmt.Print("Enter Hours Worked: ")
				fmt.Scanln(&hoursWorked)

				emp := employees.PartTimeEmployee{
					ID:          id,
					Name:        name,
					HourlyRate:  hourlyRate,
					HoursWorked: hoursWorked,
				}
				company.AddEmployee(emp)

			default:
				fmt.Println("Invalid employee type.")
			}

		case "list":
			company.ListEmployees()

		case "exit":
			fmt.Println("Exiting program...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
