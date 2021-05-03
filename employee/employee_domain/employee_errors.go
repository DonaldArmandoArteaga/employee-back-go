package employee_domain

import "errors"

const EMPLOYEE_NOT_FOUND = "employee not found"

var ErrNotFound = errors.New(EMPLOYEE_NOT_FOUND)
