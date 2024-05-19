package employee_controller

import (
	"context"
	"net/http"
)

type deleteBodyReq struct {
	EmployeeId int `json:"employeeId"`
}

// GetByCompany godoc
// @Tags         Employee
// @Description  Remove employee by ID
// @Accept       json
// @Produce      json
// @Param        Employee body deleteBodyReq true "employee data"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /employee/delete [delete]
func (c *employeeController) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var reqBody deleteBodyReq
	err := parseBodyRequest(r, &reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = c.employeeService.Delete(ctx, reqBody.EmployeeId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
