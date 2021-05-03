package employee_domain

type EmployeeRepository interface {
	GetData() ([]*Employee, error)
	GetDataById(id int64) (*Employee, error)
}
