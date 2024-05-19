package employee_repo

import (
	"context"

	"employees/internal/entities"
	queries "employees/internal/repositories/employee/queries"

	"github.com/jackc/pgx/v5/pgtype"
)

func (r *employeeRepository) GetByCompanyDepartment(
	ctx context.Context,
	companyId int,
	departmentName string,
) ([]*entities.Employee, error) {
	q := r.getQuery(ctx)
	rows, err := q.GetEmployeesByCompanyDepartment(ctx, queries.GetEmployeesByCompanyDepartmentParams{
		CompanyID: pgtype.Int4{Int32: int32(companyId), Valid: true},
		Name:      pgtype.Text{String: departmentName, Valid: true},
	})
	if err != nil {
		return nil, err
	}

	employees := make([]*entities.Employee, 0, len(rows))
	for _, v := range rows {
		employees = append(employees, repackByCompanyDepartmentEmployee(&v))
	}

	return employees, nil
}

func repackByCompanyDepartmentEmployee(row *queries.GetEmployeesByCompanyDepartmentRow) *entities.Employee {
	return &entities.Employee{
		Id:        int(row.ID),
		Name:      row.Name.String,
		Surname:   row.Surname.String,
		Phone:     row.Phone.String,
		CompanyId: int(row.CompanyID.Int32),
		Passport: entities.Passport{
			Type:   row.PassportType.String,
			Number: row.PassportNumber.String,
		},
		Department: entities.Department{
			Name:  row.DepartmentName.String,
			Phone: row.DepartmentPhone.String,
		},
	}
}
