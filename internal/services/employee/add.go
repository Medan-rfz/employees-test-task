package employee_service

import (
	"context"

	"employees/internal/entities"
)

func (s *employeeService) Add(ctx context.Context, employee *entities.Employee) (int, error) {
	var employeeId int

	err := s.txExecutor.TxBegin(ctx, func(ctx context.Context) error {
		departmentId, err := s.employeeRepo.CreateDepartment(ctx, employee.Department)
		if err != nil {
			return err
		}

		passportId, err := s.employeeRepo.CreatePassport(ctx, employee.Passport)
		if err != nil {
			return err
		}

		employeeId, err = s.employeeRepo.CreateEmployee(ctx, &entities.EmployeeDB{
			Id:           employee.Id,
			Name:         employee.Name,
			Surname:      employee.Surname,
			Phone:        employee.Phone,
			CompanyId:    employee.CompanyId,
			PassportId:   passportId,
			DepartmentId: departmentId,
		})
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return employeeId, nil
}
