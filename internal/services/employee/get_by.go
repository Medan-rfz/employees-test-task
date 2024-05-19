package employee_service

import (
	"context"

	"employees/internal/entities"
)

func (s *employeeService) GetByCompany(ctx context.Context, companyId int) ([]*entities.Employee, error) {
	return s.employeeRepo.GetByCompany(ctx, companyId)
}

func (s *employeeService) GetByCompanyDepartment(ctx context.Context, companyId int, departmentName string) ([]*entities.Employee, error) {
	return s.employeeRepo.GetByCompanyDepartment(ctx, companyId, departmentName)
}
