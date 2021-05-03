package main

import (
	"os"

	EmployeeDomain "github.com/DonaldArmandoArteaga/go-rest-api/employee/employee_domain"
	EmployeeExternalRequest "github.com/DonaldArmandoArteaga/go-rest-api/employee/external_request"
	EmployeeController "github.com/DonaldArmandoArteaga/go-rest-api/employee/rest_controller"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	EmployeeController.CreateemployeeController(
		&EmployeeController.EmployeeController{
			Gin: r,
			EmployeeServices: &EmployeeDomain.EmployeeService{
				EmployeeRepository: EmployeeExternalRequest.CreateEmployeeRepository(),
			},
		},
	)

	os.Setenv("ServiceURL", "http://masglobaltestapi.azurewebsites.net/api/Employees")
	r.Run()

}
