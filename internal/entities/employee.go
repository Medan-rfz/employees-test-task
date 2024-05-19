package entities

type Employee struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Surname   string `json:"surname"`
	Phone     string `json:"phone"`
	CompanyId int    `json:"companyId"`

	Passport   Passport   `json:"passport"`
	Department Department `json:"department"`
}

type EmployeeDB struct {
	Id        int
	Name      string
	Surname   string
	Phone     string
	CompanyId int

	PassportId   int
	DepartmentId int
}
