package employee_repo

import (
	"context"

	"employees/internal/entities"
	queries "employees/internal/repositories/employee/queries"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *employeeRepository) UpdateEmployee(ctx context.Context, employee *entities.EmployeeDB) error {
	q := r.getQuery(ctx)
	return q.UpdateEmployee(ctx, queries.UpdateEmployeeParams{
		ID:        int32(employee.Id),
		Name:      pgtype.Text{String: employee.Name, Valid: true},
		Surname:   pgtype.Text{String: employee.Surname, Valid: true},
		Phone:     pgtype.Text{String: employee.Phone, Valid: true},
		CompanyID: pgtype.Int4{Int32: int32(employee.CompanyId), Valid: true},
	})
}

func (r *employeeRepository) UpdateDepartment(ctx context.Context, id int, department *entities.Department) error {
	q := r.getQuery(ctx)
	return q.UpdateDepartment(ctx, queries.UpdateDepartmentParams{
		ID:    int32(id),
		Name:  pgtype.Text{String: department.Name, Valid: true},
		Phone: pgtype.Text{String: department.Phone, Valid: true},
	})
}

func (r *employeeRepository) UpdatePassport(ctx context.Context, id int, passport *entities.Passport) error {
	q := r.getQuery(ctx)
	return q.UpdatePassport(ctx, queries.UpdatePassportParams{
		ID:     int32(id),
		Number: pgtype.Text{String: passport.Number, Valid: true},
		Type:   pgtype.Text{String: passport.Type, Valid: true},
	})
}
