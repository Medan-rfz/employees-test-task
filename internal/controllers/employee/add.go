package employee_controller

import (
	"context"
	"encoding/json"
	"net/http"

	"employees/internal/entities"
)

type addBodyReq struct {
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Phone     string `json:"phone"`
	CompanyId int    `json:"companyId"`

	Passport   entities.Passport   `json:"passport"`
	Department entities.Department `json:"department"`
}

type addBodyResp struct {
	EmployeeId int `json:"employeeId"`
}

// Add godoc
// @Tags         Employee
// @Description  Add employee
// @Accept       json
// @Produce      json
// @Param        Employee body addBodyReq true "employee data"
// @Success      200 {object} addBodyResp
// @Failure      400
// @Failure      500
// @Router       /employee/add [post]
func (c *employeeController) Add(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var reqBody addBodyReq
	err := parseBodyRequest(r, &reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	employeeId, err := c.employeeService.Add(ctx, createEmployee(reqBody))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respBody, err := json.Marshal(addBodyResp{
		EmployeeId: employeeId,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(respBody)
}

func createEmployee(dto addBodyReq) *entities.Employee {
	return &entities.Employee{
		Name:       dto.Name,
		Surname:    dto.Surname,
		Phone:      dto.Phone,
		CompanyId:  dto.CompanyId,
		Passport:   dto.Passport,
		Department: dto.Department,
	}
}
