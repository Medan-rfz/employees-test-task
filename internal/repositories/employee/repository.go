package employee_repo

import (
	"context"

	queries "employees/internal/repositories/employee/queries"
	txexecutor "employees/internal/repositories/tx_executor"

	"github.com/jackc/pgx/v5"
)

type employeeRepository struct {
	conn    *pgx.Conn
	queries *queries.Queries
}

func NewRepository(conn *pgx.Conn) *employeeRepository {
	return &employeeRepository{
		conn:    conn,
		queries: queries.New(conn),
	}
}

func (r *employeeRepository) getQuery(ctx context.Context) *queries.Queries {
	q := r.queries
	if txTemp := ctx.Value(txexecutor.TxContextMetaName); txTemp != nil {
		if tx, ok := txTemp.(pgx.Tx); ok {
			q = q.WithTx(tx)
		}
	}

	return q
}
