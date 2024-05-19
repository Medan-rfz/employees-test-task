package employee_service

import (
	"context"

	"employees/internal/entities"
)

type txExecutor interface {
	TxBegin(ctx context.Context, exec func(ctx context.Context) error) error
}

type employeeRepo interface {
	CreateEmployee(ctx context.Context, employee *entities.EmployeeDB) (int, error)
	CreateDepartment(ctx context.Context, department entities.Department) (int, error)
	CreatePassport(ctx context.Context, passport entities.Passport) (int, error)

	UpdateEmployee(ctx context.Context, employee *entities.EmployeeDB) error
	UpdateDepartment(ctx context.Context, id int, department *entities.Department) error
	UpdatePassport(ctx context.Context, id int, passport *entities.Passport) error

	Delete(ctx context.Context, employeeId int) error

	GetEmployeeById(ctx context.Context, id int) (*entities.EmployeeDB, error)
	GetDepartmentById(ctx context.Context, id int) (*entities.Department, error)
	GetPassportById(ctx context.Context, id int) (*entities.Passport, error)
	GetByCompany(ctx context.Context, companyId int) ([]*entities.Employee, error)
	GetByCompanyDepartment(ctx context.Context, companyId int, departmentName string) ([]*entities.Employee, error)
}

type employeeService struct {
	txExecutor   txExecutor
	employeeRepo employeeRepo
}

func NewService(txExecutor txExecutor, employeeRepo employeeRepo) *employeeService {
	return &employeeService{
		txExecutor:   txExecutor,
		employeeRepo: employeeRepo,
	}
}
