package employee_controller

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"employees/internal/dto"
	"employees/internal/entities"
)

type employeeService interface {
	Add(ctx context.Context, employee *entities.Employee) (int, error)
	Delete(ctx context.Context, employeeId int) error
	Update(ctx context.Context, updateFields dto.EmployeeDTO) error
	GetByCompany(ctx context.Context, companyId int) ([]*entities.Employee, error)
	GetByCompanyDepartment(ctx context.Context, companyId int, departmentName string) ([]*entities.Employee, error)
}

type employeeController struct {
	employeeService employeeService
}

func NewController(employeeService employeeService) *employeeController {
	return &employeeController{
		employeeService: employeeService,
	}
}

func parseBodyRequest[T any](r *http.Request, obj *T) error {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, obj); err != nil {
		return err
	}

	// TODO Add validation

	return nil
}
