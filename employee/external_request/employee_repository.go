package employee

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	EmployeeDomain "github.com/DonaldArmandoArteaga/go-rest-api/employee/employee_domain"
)

type EmployeeRepository struct {
}

func CreateEmployeeRepository() *EmployeeRepository {
	return &EmployeeRepository{}
}

func (ar *EmployeeRepository) GetData() ([]*EmployeeDomain.Employee, error) {

	response, err := http.Get(os.Getenv("ServiceURL"))

	if err != nil {
		return []*EmployeeDomain.Employee{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return []*EmployeeDomain.Employee{}, err
	}

	var employeeModelRepository []*EmployeeModelRepository
	json.Unmarshal(responseData, &employeeModelRepository)

	employees := []*EmployeeDomain.Employee{}

	for _, employee := range employeeModelRepository {
		e, err := EmployeeModelRepositoryToemployee(employee)

		if err != nil {
			return []*EmployeeDomain.Employee{}, err
		}

		employees = append(employees, e)
	}

	return employees, nil
}

func (ar *EmployeeRepository) GetDataById(id int64) (*EmployeeDomain.Employee, error) {
	response, err := http.Get(os.Getenv("ServiceURL"))

	if err != nil {
		return &EmployeeDomain.Employee{}, err
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return &EmployeeDomain.Employee{}, err
	}

	var employeeModelRepository []*EmployeeModelRepository
	json.Unmarshal(responseData, &employeeModelRepository)

	for _, employee := range employeeModelRepository {

		if employee.Id == id {
			e, err := EmployeeModelRepositoryToemployee(employee)

			if err != nil {
				return &EmployeeDomain.Employee{}, err
			}

			return e, nil
		}

	}
	return &EmployeeDomain.Employee{}, nil
}
