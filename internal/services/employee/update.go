package employee_service

import (
	"context"

	"employees/internal/dto"
	"employees/internal/entities"
)

func (s *employeeService) Update(ctx context.Context, newEmployee dto.EmployeeDTO) error {
	if newEmployee.Id == nil {
		return entities.ErrEmployeeId
	}

	return s.txExecutor.TxBegin(ctx, func(ctx context.Context) error {
		employee, err := s.employeeRepo.GetEmployeeById(ctx, *newEmployee.Id)
		if err != nil {
			return entities.ErrEmployeeNotFound
		}

		department, err := s.employeeRepo.GetDepartmentById(ctx, employee.DepartmentId)
		if err != nil {
			return err
		}

		passport, err := s.employeeRepo.GetPassportById(ctx, employee.PassportId)
		if err != nil {
			return err
		}

		updateFieldsEmployee(newEmployee, employee)
		updateFieldsDepartment(newEmployee.Department, department)
		updateFieldsPassport(newEmployee.Passport, passport)

		if err := s.employeeRepo.UpdateEmployee(ctx, employee); err != nil {
			return err
		}

		if err := s.employeeRepo.UpdateDepartment(ctx, employee.DepartmentId, department); err != nil {
			return err
		}

		if err := s.employeeRepo.UpdatePassport(ctx, employee.PassportId, passport); err != nil {
			return err
		}

		return nil
	})
}

func updateFieldsEmployee(new dto.EmployeeDTO, employee *entities.EmployeeDB) {
	if new.Name != nil {
		employee.Name = *new.Name
	}
	if new.Surname != nil {
		employee.Surname = *new.Surname
	}
	if new.Phone != nil {
		employee.Phone = *new.Phone
	}
	if new.CompanyId != nil {
		employee.CompanyId = *new.CompanyId
	}
}

func updateFieldsDepartment(new *dto.DepartmentDTO, department *entities.Department) {
	if new == nil {
		return
	}

	if new.Name != nil {
		department.Name = *new.Name
	}
	if new.Phone != nil {
		department.Phone = *new.Phone
	}
}

func updateFieldsPassport(new *dto.PassportDTO, passport *entities.Passport) {
	if new == nil {
		return
	}

	if new.Number != nil {
		passport.Number = *new.Number
	}
	if new.Type != nil {
		passport.Type = *new.Type
	}
}
