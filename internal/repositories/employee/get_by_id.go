package employee_repo

import (
	"context"

	"employees/internal/entities"
	queries "employees/internal/repositories/employee/queries"
)

func (r *employeeRepository) GetEmployeeById(
	ctx context.Context,
	id int,
) (*entities.EmployeeDB, error) {
	q := r.getQuery(ctx)
	row, err := q.GetEmployeeById(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return repackGetByIdEmployee(&row), nil
}

func (r *employeeRepository) GetDepartmentById(
	ctx context.Context,
	id int,
) (*entities.Department, error) {
	q := r.getQuery(ctx)
	row, err := q.GetDepartmentById(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return repackGetByIdDepartment(&row), nil
}

func (r *employeeRepository) GetPassportById(
	ctx context.Context,
	id int,
) (*entities.Passport, error) {
	q := r.getQuery(ctx)
	row, err := q.GetPassportById(ctx, int32(id))
	if err != nil {
		return nil, err
	}

	return repackGetByIdPassport(&row), nil
}

func repackGetByIdEmployee(row *queries.Employee) *entities.EmployeeDB {
	return &entities.EmployeeDB{
		Id:           int(row.ID),
		Name:         row.Name.String,
		Surname:      row.Surname.String,
		Phone:        row.Phone.String,
		CompanyId:    int(row.CompanyID.Int32),
		PassportId:   int(row.PassportID.Int32),
		DepartmentId: int(row.DepartmentID.Int32),
	}
}

func repackGetByIdPassport(row *queries.GetPassportByIdRow) *entities.Passport {
	return &entities.Passport{
		Type:   row.Type.String,
		Number: row.Number.String,
	}
}

func repackGetByIdDepartment(row *queries.GetDepartmentByIdRow) *entities.Department {
	return &entities.Department{
		Name:  row.Name.String,
		Phone: row.Phone.String,
	}
}
