package employee_repo

import (
	"context"

	"employees/internal/entities"
	queries "employees/internal/repositories/employee/queries"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *employeeRepository) CreatePassport(ctx context.Context, passport entities.Passport) (int, error) {
	q := r.getQuery(ctx)
	id, err := q.CreatePassport(ctx, queries.CreatePassportParams{
		Type:   pgtype.Text{String: passport.Type, Valid: true},
		Number: pgtype.Text{String: passport.Number, Valid: true},
	})

	return int(id), err
}
