package employee_domain

import "fmt"

type EmployeeServices interface {
	GetAll() ([]*Employee, error)
	GetById(id int64) (*Employee, error)
}

type EmployeeService struct {
	EmployeeRepository EmployeeRepository
}

func (as *EmployeeService) GetAll() ([]*Employee, error) {

	employee, err := as.EmployeeRepository.GetData()

	if err != nil {
		return []*Employee{}, err
	}

	for i := range employee {

		var errAnualSalary error
		employee[i].annualSalary, errAnualSalary = annualSalaryFormula(
			employee[i].ContractTypeName(),
			employee[i].HourlySalary(),
			employee[i].MonthlySalary(),
		)

		if errAnualSalary != nil {
			return []*Employee{}, errAnualSalary
		}

	}

	return employee, nil
}

func (as *EmployeeService) GetById(id int64) (*Employee, error) {

	employee, err := as.EmployeeRepository.GetDataById(id)

	if err != nil {
		return employee, err
	}

	if *employee == (Employee{}) {
		return employee, ErrNotFound
	}

	var errAnualSalary error
	employee.annualSalary, errAnualSalary = annualSalaryFormula(
		employee.ContractTypeName(),
		employee.HourlySalary(),
		employee.MonthlySalary(),
	)

	if errAnualSalary != nil {
		return &Employee{}, errAnualSalary
	}

	return employee, nil
}

func annualSalaryFormula(
	contractTypeName string,
	hourlySalary float64,
	monthlySalary float64,
) (float64, error) {

	if contractTypeName == HOURLY_SALARY {
		return (hourlySalary * 120 * 12), nil
	}

	if contractTypeName == MONTHLY_SALARY {
		return (monthlySalary * 12), nil

	}

	return 0, fmt.Errorf("error when none of the allowed values arrive")

}
