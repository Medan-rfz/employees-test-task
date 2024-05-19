package employee_controller

import (
	"context"
	"net/http"

	"employees/internal/dto"
)

// type updateBodyReq struct {
// 	Id        *int    `json:"id"`
// 	Name      *string `json:"name"`
// 	Surname   *string `json:"surname"`
// 	Phone     *string `json:"phone"`
// 	CompanyId *int    `json:"companyId"`

// 	Passport   *entities.Passport   `json:"passport"`
// 	Department *entities.Department `json:"department"`
// }

// Update godoc
// @Tags         Employee
// @Description  Update employee data
// @Accept       json
// @Produce      json
// @Param        Employee body dto.EmployeeDTO true "employee data"
// @Success      200
// @Failure      400
// @Failure      500
// @Router       /employee/update [post]
func (c *employeeController) Update(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	var reqBody dto.EmployeeDTO
	err := parseBodyRequest(r, &reqBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := c.employeeService.Update(ctx, reqBody); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}
