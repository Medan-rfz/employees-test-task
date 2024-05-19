package employee_repo

import (
	"context"

	"employees/internal/entities"
	queries "employees/internal/repositories/employee/queries"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *employeeRepository) CreateDepartment(ctx context.Context, department entities.Department) (int, error) {
	q := r.getQuery(ctx)
	id, err := q.GetOrCreateDepartment(ctx, queries.GetOrCreateDepartmentParams{
		Name:  pgtype.Text{String: department.Name, Valid: true},
		Phone: pgtype.Text{String: department.Phone, Valid: true},
	})

	return int(id), err
}
