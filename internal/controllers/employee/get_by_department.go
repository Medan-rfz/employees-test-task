package employee_controller

import (
	"context"
	"encoding/json"
	"net/http"

	"employees/internal/entities"
)

type getByCompanyDepartmentBodyReq struct {
	CompanyId      int    `json:"companyId"`
	DepartmentName string `json:"departmentName"`
}

type getByCompanyDepartmentBodyResp struct {
	Employees []*entities.Employee `json:"employees"`
}

// GetByCompanyDepartment godoc
// @Tags         Employee
// @Description  Get employee by company ID and department name
// @Accept       json
// @Produce      json
// @Param        Employee body getByCompanyDepartmentBodyReq true "employee data"
// @Success      200 {object} getByCompanyDepartmentBodyResp
// @Failure      400
// @Failure      500
// @Router       /employee/get/by_company_department [post]
func (c *employeeController) GetByCompanyDepartment(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var reqBody getByCompanyDepartmentBodyReq
	err := parseBodyRequest(r, &reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	employees, err := c.employeeService.GetByCompanyDepartment(ctx, reqBody.CompanyId, reqBody.DepartmentName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(getByCompanyDepartmentBodyResp{
		Employees: employees,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(respBody)
}
