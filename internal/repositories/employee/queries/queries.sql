-- name: CreateEmployee :one
INSERT INTO employee
    (name, surname, phone, company_id, department_id, passport_id)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING id;

-- name: GetOrCreateDepartment :one
WITH existing_department AS (
    SELECT d.id
    FROM department d
    WHERE d.name = $1 AND d.phone = $2
),
new_department AS (
    INSERT INTO department (name, phone)
    SELECT $1, $2
    WHERE NOT EXISTS (SELECT id FROM existing_department)
    RETURNING id
)
SELECT id FROM existing_department
UNION ALL
SELECT id FROM new_department;

-- name: CreatePassport :one
INSERT INTO passport
    (number, type)
VALUES
    ($1, $2)
RETURNING id;

-- name: UpdateEmployee :exec
UPDATE employee
SET
    name = $2,
    surname = $3,
    phone = $4,
    company_id = $5
WHERE id = $1;

-- name: UpdateDepartment :exec
UPDATE department
SET
    name = $2,
    phone = $3
WHERE id = $1;

-- name: UpdatePassport :exec
UPDATE passport
SET
    number = $2,
    type = $3
WHERE id = $1;

-- name: DeleteEmployee :exec
DELETE FROM employee e
WHERE e.id = $1;

-- name: GetEmployeeById :one
SELECT 
    e.id, e.name, e.surname, e.phone, e.company_id, e.passport_id, e.department_id
FROM employee e
WHERE e.id = $1;

-- name: GetDepartmentById :one
SELECT 
    d.name, d.phone
FROM department d
WHERE d.id = $1;

-- name: GetPassportById :one
SELECT 
    p.number, p.type
FROM passport p
WHERE p.id = $1;

-- name: GetEmployeesByCompany :many
SELECT 
    e.id, e.name, e.surname, e.phone, e.company_id,
    d.name AS department_name, 
    d.phone AS department_phone,
    p.number AS passport_number, 
    p.type AS passport_type
FROM employee e
JOIN department d ON d.id = e.department_id
JOIN passport p ON p.id = e.passport_id
WHERE e.company_id = $1;

-- name: GetEmployeesByCompanyDepartment :many
SELECT 
    e.id, e.name, e.surname, e.phone, e.company_id,
    d.name AS department_name, 
    d.phone AS department_phone,
    p.number AS passport_number, 
    p.type AS passport_type
FROM employee e
JOIN department d ON d.id = e.department_id
JOIN passport p ON p.id = e.passport_id
WHERE e.company_id = $1 AND d.name = $2;
