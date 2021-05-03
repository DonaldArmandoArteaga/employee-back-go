package employee

import (
	EmployeeDomain "github.com/DonaldArmandoArteaga/go-rest-api/employee/employee_domain"
)

type EmployeeModelRepository struct {
	Id               int64
	Name             string
	ContractTypeName string
	RoleName         string
	RoleDescription  string
	HourlySalary     float64
	MonthlySalary    float64
}

func EmployeeModelRepositoryToemployee(e *EmployeeModelRepository) (*EmployeeDomain.Employee, error) {

	employee, err := EmployeeDomain.Createemployee(
		&EmployeeDomain.CreateEmployeeImput{
			Id:               e.Id,
			Name:             e.Name,
			ContractTypeName: e.ContractTypeName,
			RoleName:         e.RoleName,
			RoleDescription:  e.RoleDescription,
			HourlySalary:     e.HourlySalary,
			MonthlySalary:    e.MonthlySalary,
		},
	)

	if err != nil {
		return &EmployeeDomain.Employee{}, err
	}

	return employee, nil
}

func EmployeeToemployeeModelRepository(ed *EmployeeDomain.Employee) *EmployeeModelRepository {

	return &EmployeeModelRepository{
		Id:               ed.Id(),
		Name:             ed.Name(),
		ContractTypeName: ed.ContractTypeName(),
		RoleName:         ed.RoleName(),
		RoleDescription:  ed.RoleDescription(),
		HourlySalary:     ed.HourlySalary(),
		MonthlySalary:    ed.HourlySalary(),
	}
}
