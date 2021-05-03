package employee

import (
	EmployeeDomain "github.com/DonaldArmandoArteaga/go-rest-api/employee/employee_domain"
)

type employeeQueryDTO struct {
	Id               int64   `json:"id"`
	Name             string  `json:"name"`
	ContractTypeName string  `json:"contractTypeName"`
	RoleName         string  `json:"roleName"`
	RoleDescription  string  `json:"roleDescription"`
	HourlySalary     float64 `json:"hourlySalary"`
	MonthlySalary    float64 `json:"monthlySalary"`
	AnnualSalary     float64 `json:"annualSalary"`
}

func EmployeeDTOToemployee(eqdto *employeeQueryDTO) (*EmployeeDomain.Employee, error) {

	employee, err := EmployeeDomain.Createemployee(
		&EmployeeDomain.CreateEmployeeImput{
			Id:               eqdto.Id,
			Name:             eqdto.Name,
			ContractTypeName: eqdto.ContractTypeName,
			RoleName:         eqdto.RoleName,
			RoleDescription:  eqdto.RoleDescription,
			HourlySalary:     eqdto.HourlySalary,
			MonthlySalary:    eqdto.MonthlySalary,
			AnnualSalary:     eqdto.AnnualSalary,
		},
	)

	if err != nil {
		return &EmployeeDomain.Employee{}, err
	}

	return employee, nil
}

func EmployeeToemployeeQueryDTO(e *EmployeeDomain.Employee) *employeeQueryDTO {
	return &employeeQueryDTO{
		Id:               e.Id(),
		Name:             e.Name(),
		ContractTypeName: e.ContractTypeName(),
		RoleName:         e.RoleName(),
		RoleDescription:  e.RoleDescription(),
		HourlySalary:     e.HourlySalary(),
		MonthlySalary:    e.MonthlySalary(),
		AnnualSalary:     e.AnnualSalary(),
	}
}
