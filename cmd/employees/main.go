package main

import (
	"context"
	"fmt"
	"log"
	"os"

	employee_controller "employees/internal/controllers/employee"
	employee_repo "employees/internal/repositories/employee"
	txexecutor "employees/internal/repositories/tx_executor"
	server_http "employees/internal/servers/http"
	employee_service "employees/internal/services/employee"

	"github.com/jackc/pgx/v5"
)

const (
	httpPort = 8081
)

func main() {
	ctx := context.Background()

	connStr, exist := os.LookupEnv("EMPLOYEES_DB_CONN_STR")
	if !exist {
		log.Fatalln("DB connection string now found in environment variables")
	}

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close(ctx)

	txExecutor, err := txexecutor.NewTxExecutor(conn)
	if err != nil {
		log.Fatalln(err)
	}

	employeeRepo := employee_repo.NewRepository(conn)
	employeeService := employee_service.NewService(txExecutor, employeeRepo)
	employeeController := employee_controller.NewController(employeeService)
	httpServer := server_http.NewHttpServer(employeeController)

	httpServer.Run(server_http.HttpServerConfig{
		Address: fmt.Sprintf(":%v", httpPort),
	})
}
