package employee

import (
	"net/http"
	"strconv"

	EmployeeDomain "github.com/DonaldArmandoArteaga/go-rest-api/employee/employee_domain"
	"github.com/gin-gonic/gin"
)

type EmployeeController struct {
	Gin              *gin.Engine
	EmployeeServices EmployeeDomain.EmployeeServices
}

func CreateemployeeController(ac *EmployeeController) {
	ac.Gin.GET("/employees", ac.getAllemployee)
	ac.Gin.GET("/employees/:id", ac.getEmployeeById)
}

func (ac *EmployeeController) getAllemployee(c *gin.Context) {

	employee, err := ac.EmployeeServices.GetAll()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	employeeDTO := []*employeeQueryDTO{}

	for _, employeeRepository := range employee {
		employeeDTO = append(employeeDTO, EmployeeToemployeeQueryDTO(employeeRepository))
	}

	c.JSON(http.StatusOK, gin.H{"employees": employeeDTO})

}

func (ac *EmployeeController) getEmployeeById(c *gin.Context) {

	id, parseIdErr := strconv.ParseInt(c.Param("id"), 10, 32)

	if parseIdErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": parseIdErr.Error()})
		return
	}

	employee, err := ac.EmployeeServices.GetById(id)

	if err != nil {
		if err.Error() == EmployeeDomain.ErrNotFound.Error() {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"employee": EmployeeToemployeeQueryDTO(employee)})
}
