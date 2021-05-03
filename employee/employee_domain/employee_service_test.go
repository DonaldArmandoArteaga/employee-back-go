package employee_domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MyMockedObject struct {
	mock.Mock
}

func (m *MyMockedObject) GetDataById(id int64) (*Employee, error) {

	args := m.Called(id)
	return args.Get(0).(*Employee), args.Error(1)

}

func (m *MyMockedObject) GetData() ([]*Employee, error) {

	args := m.Called()
	return args.Get(0).([]*Employee), args.Error(1)

}

func TestListByIdEmployeeWhenHoulrySalaryInContractType(t *testing.T) {

	e := &CreateEmployeeImput{
		Id:               1,
		Name:             "any name",
		ContractTypeName: HOURLY_SALARY,
		RoleName:         "any role",
		RoleDescription:  "",
		HourlySalary:     60000.0,
		MonthlySalary:    80000.0,
		AnnualSalary:     0,
	}

	employeeWithoutAnnualSalary, _ := Createemployee(e)

	e.AnnualSalary = 86400000
	employeeWithAnnualSalary, _ := Createemployee(e)

	testObj := new(MyMockedObject)

	var ii int64 = 1
	testObj.On("GetDataById", ii).Return(employeeWithoutAnnualSalary, nil)

	ee := EmployeeService{
		EmployeeRepository: testObj,
	}

	valueExpected, _ := ee.GetById(ii)

	assert.Equal(t, valueExpected, employeeWithAnnualSalary)

}

func TestListByIdEmployeeWhenIsEmpty(t *testing.T) {

	var ii int64 = 3

	testObj := new(MyMockedObject)
	testObj.On("GetDataById", ii).Return(&Employee{}, nil)

	ee := EmployeeService{
		EmployeeRepository: testObj,
	}

	_, err := ee.GetById(ii)

	assert.EqualError(t, err, EMPLOYEE_NOT_FOUND)
}

func TestListByIdEmployeeWhenMonthlySalaryInContractType(t *testing.T) {
	e := &CreateEmployeeImput{
		Id:               1,
		Name:             "any name",
		ContractTypeName: MONTHLY_SALARY,
		RoleName:         "any role",
		RoleDescription:  "",
		HourlySalary:     60000.0,
		MonthlySalary:    80000.0,
		AnnualSalary:     0,
	}

	employeeWithoutAnnualSalary, _ := Createemployee(e)

	e.AnnualSalary = 960000
	employeeWithAnnualSalary, _ := Createemployee(e)

	testObj := new(MyMockedObject)

	var ii int64 = 1
	testObj.On("GetDataById", ii).Return(employeeWithoutAnnualSalary, nil)

	ee := EmployeeService{
		EmployeeRepository: testObj,
	}

	valueExpected, _ := ee.GetById(ii)

	assert.Equal(t, valueExpected, employeeWithAnnualSalary)
}

func TestLlistByIdEmployeeWhenIsEmpty(t *testing.T) {
	testObj := new(MyMockedObject)

	testObj.On("GetData").Return([]*Employee{}, nil)

	ee := EmployeeService{
		EmployeeRepository: testObj,
	}

	valueExpected, _ := ee.GetAll()

	assert.Equal(t, valueExpected, []*Employee{})
}

func TestListEmployee(t *testing.T) {

	ehs := &CreateEmployeeImput{
		Id:               1,
		Name:             "any name",
		ContractTypeName: HOURLY_SALARY,
		RoleName:         "any role",
		RoleDescription:  "",
		HourlySalary:     60000.0,
		MonthlySalary:    80000.0,
		AnnualSalary:     0,
	}

	employeeWithoutAnnualSalaryhs, _ := Createemployee(ehs)

	ehs.AnnualSalary = 86400000
	employeeWithAnnualSalaryhs, _ := Createemployee(ehs)

	ems := &CreateEmployeeImput{
		Id:               1,
		Name:             "any name",
		ContractTypeName: MONTHLY_SALARY,
		RoleName:         "any role",
		RoleDescription:  "",
		HourlySalary:     60000.0,
		MonthlySalary:    80000.0,
		AnnualSalary:     0,
	}

	employeeWithoutAnnualSalaryms, _ := Createemployee(ems)

	ems.AnnualSalary = 960000
	employeeWithAnnualSalaryms, _ := Createemployee(ems)

	testObj := new(MyMockedObject)

	valueShouldBe := []*Employee{employeeWithoutAnnualSalaryhs, employeeWithoutAnnualSalaryms}
	testObj.On("GetData").Return(valueShouldBe, nil)

	ee := EmployeeService{
		EmployeeRepository: testObj,
	}

	valueExpected, _ := ee.GetAll()

	assert.Equal(t, *valueExpected[0], *employeeWithAnnualSalaryhs)
	assert.Equal(t, *valueExpected[1], *employeeWithAnnualSalaryms)

}
