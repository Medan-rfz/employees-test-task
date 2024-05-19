package employee_repo

import "context"

func (r *employeeRepository) Delete(ctx context.Context, employeeId int) error {
	q := r.getQuery(ctx)
	return q.DeleteEmployee(ctx, int32(employeeId))
}
