package employee_repo

import (
	"context"

	"employees/internal/entities"
	queries "employees/internal/repositories/employee/queries"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *employeeRepository) CreateEmployee(ctx context.Context, employee *entities.EmployeeDB) (int, error) {
	q := r.getQuery(ctx)
	id, err := q.CreateEmployee(ctx, queries.CreateEmployeeParams{
		Name:         pgtype.Text{String: employee.Name, Valid: true},
		Surname:      pgtype.Text{String: employee.Surname, Valid: true},
		Phone:        pgtype.Text{String: employee.Phone, Valid: true},
		CompanyID:    pgtype.Int4{Int32: int32(employee.CompanyId), Valid: true},
		DepartmentID: pgtype.Int4{Int32: int32(employee.DepartmentId), Valid: true},
		PassportID:   pgtype.Int4{Int32: int32(employee.PassportId), Valid: true},
	})

	return int(id), err
}
