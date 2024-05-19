
build:
	go build -o ./bin/ ./cmd/employees

run:
	go run ./cmd/employees

############################################
# Sqlc
############################################ 
sqlc-gen:
	cd ./sqlc && sqlc generate

############################################
# Swagger
############################################ 
swagger-gen:
	swag init -g server.go -d "./internal/servers/http,./internal/controllers/employee,./internal/entities,./internal/dto"

############################################
# Migrations
############################################ 
migration-up:
	goose -dir migrations postgres ${EMPLOYEES_DB_CONN_STR} up

migration-down:
	goose -dir migrations postgres ${EMPLOYEES_DB_CONN_STR} down
