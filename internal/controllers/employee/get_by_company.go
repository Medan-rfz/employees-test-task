package employee_controller

import (
	"context"
	"encoding/json"
	"net/http"

	"employees/internal/entities"
)

type getByCompanyBodyReq struct {
	CompanyId int `json:"companyId"`
}

type getByCompanyBodyResp struct {
	Employees []*entities.Employee `json:"employees"`
}

// GetByCompany godoc
// @Tags         Employee
// @Description  Get employee by company ID
// @Accept       json
// @Produce      json
// @Param        Employee body getByCompanyBodyReq true "employee data"
// @Success      200 {object} getByCompanyBodyResp
// @Failure      400
// @Failure      500
// @Router       /employee/get/by_company [post]
func (c *employeeController) GetByCompany(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var reqBody getByCompanyBodyReq
	err := parseBodyRequest(r, &reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	employees, err := c.employeeService.GetByCompany(ctx, reqBody.CompanyId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(getByCompanyBodyResp{
		Employees: employees,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(respBody)
}
