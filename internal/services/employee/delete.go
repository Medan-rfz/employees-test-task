package employee_service

import "context"

func (s *employeeService) Delete(ctx context.Context, employeeId int) error {
	return s.employeeRepo.Delete(ctx, employeeId)
}
