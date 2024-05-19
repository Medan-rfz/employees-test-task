-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "department" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    phone VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS "passport" (
    id SERIAL PRIMARY KEY,
    number VARCHAR(50),
    type VARCHAR(50)
);

CREATE TABLE IF NOT EXISTS "employee" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50),
    surname VARCHAR(50),
    phone VARCHAR(50),
    company_id SERIAL,

    passport_id SERIAL REFERENCES "passport"(id),
    department_id SERIAL REFERENCES "department"(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "employee" CASCADE;
DROP TABLE IF EXISTS "department" CASCADE;
DROP TABLE IF EXISTS "passport" CASCADE;
-- +goose StatementEnd
