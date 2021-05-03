package employee_domain

import (
	"fmt"
	"strings"
)

const (
	HOURLY_SALARY  = "HourlySalaryEmployee"
	MONTHLY_SALARY = "MonthlySalaryEmployee"
)

type (
	CreateEmployeeImput struct {
		Id               int64
		Name             string
		ContractTypeName string
		RoleName         string
		RoleDescription  string
		HourlySalary     float64
		MonthlySalary    float64
		AnnualSalary     float64
	}

	Employee struct {
		id               int64
		name             string
		contractTypeName string
		roleName         string
		roleDescription  string
		hourlySalary     float64
		monthlySalary    float64
		annualSalary     float64
	}
)

func Createemployee(cai *CreateEmployeeImput) (*Employee, error) {

	createemployeeErrors := []string{}

	if cai.Id < 0 {
		createemployeeErrors = append(
			createemployeeErrors,
			"id must be positive",
		)
	}

	if cai.Name == "" {
		createemployeeErrors = append(
			createemployeeErrors,
			"Name is required",
		)
	}

	if cai.ContractTypeName == "" {
		createemployeeErrors = append(
			createemployeeErrors,
			"Contract Type Name is required",
		)
	}

	if cai.RoleName == "" {
		createemployeeErrors = append(
			createemployeeErrors,
			"Role Name is required",
		)
	}

	if cai.HourlySalary < 0 {
		createemployeeErrors = append(
			createemployeeErrors,
			"Hourly Salary must be positive",
		)
	}

	if cai.MonthlySalary < 0 {
		createemployeeErrors = append(
			createemployeeErrors,
			"Monthly Salary must be positive",
		)
	}

	if !(cai.ContractTypeName == HOURLY_SALARY || cai.ContractTypeName == MONTHLY_SALARY) {
		createemployeeErrors = append(
			createemployeeErrors,
			"Error when none of the allowed values arrive in Contract Type Name",
		)
	}

	if len(createemployeeErrors) > 0 {
		return &Employee{}, fmt.Errorf(strings.Join(createemployeeErrors, "\n"))
	}

	return &Employee{
		id:               cai.Id,
		name:             cai.Name,
		contractTypeName: cai.ContractTypeName,
		roleName:         cai.RoleName,
		roleDescription:  cai.RoleDescription,
		hourlySalary:     cai.HourlySalary,
		monthlySalary:    cai.MonthlySalary,
		annualSalary:     cai.AnnualSalary,
	}, nil
}

func (e *Employee) Id() int64 {
	return e.id
}

func (e *Employee) Name() string {
	return e.name
}

func (e *Employee) ContractTypeName() string {
	return e.contractTypeName
}

func (e *Employee) RoleName() string {
	return e.roleName
}

func (e *Employee) RoleDescription() string {
	return e.roleDescription
}

func (e *Employee) HourlySalary() float64 {
	return e.hourlySalary
}

func (e *Employee) MonthlySalary() float64 {
	return e.monthlySalary
}

func (e *Employee) AnnualSalary() float64 {
	return e.annualSalary
}
