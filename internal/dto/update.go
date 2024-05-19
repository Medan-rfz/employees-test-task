package dto

type EmployeeDTO struct {
	Id        *int    `json:"id"`
	Name      *string `json:"name"`
	Surname   *string `json:"surname"`
	Phone     *string `json:"phone"`
	CompanyId *int    `json:"companyId"`

	Passport   *PassportDTO   `json:"passport"`
	Department *DepartmentDTO `json:"department"`
}

type PassportDTO struct {
	Type   *string `json:"type"`
	Number *string `json:"number"`
}

type DepartmentDTO struct {
	Name  *string `json:"name"`
	Phone *string `json:"phone"`
}
