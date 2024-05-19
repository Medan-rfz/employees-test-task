package server_http

import (
	"log"
	"net"
	"net/http"
	"time"

	mw "employees/internal/servers/middleware"

	"employees/docs"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"
	httpSwagger "github.com/swaggo/http-swagger"
)

type employeeController interface {
	Add(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	GetByCompany(w http.ResponseWriter, r *http.Request)
	GetByCompanyDepartment(w http.ResponseWriter, r *http.Request)
}

type httpServer struct {
	employeeController employeeController
}

type HttpServerConfig struct {
	Address string
}

func NewHttpServer(employeeController employeeController) *httpServer {
	return &httpServer{
		employeeController: employeeController,
	}
}

func (s *httpServer) Run(config HttpServerConfig) error {
	r := chi.NewRouter()
	r.Use(mw.LoggerMiddleware)
	r.Use(middleware.Timeout(time.Second * 20))
	r.Post("/employee/add", s.employeeController.Add)
	r.Delete("/employee/delete", s.employeeController.Delete)
	r.Post("/employee/update", s.employeeController.Update)
	r.Post("/employee/get/by_company", s.employeeController.GetByCompany)
	r.Post("/employee/get/by_company_department", s.employeeController.GetByCompanyDepartment)

	docs.SwaggerInfo.Title = "Employee service API"
	docs.SwaggerInfo.Description = "Employee service"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = ""
	docs.SwaggerInfo.Schemes = []string{"http"}

	r.Get(
		"/swagger/*",
		httpSwagger.Handler(httpSwagger.URL("./swagger/doc.json")),
	)

	c := cors.New(cors.Options{
		AllowedOrigins:      []string{"*"},
		AllowedMethods:      []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:      []string{"*"},
		AllowCredentials:    true,
		AllowPrivateNetwork: true,
	})
	corsMux := c.Handler(r)

	conn, err := net.Listen("tcp", config.Address)
	if err != nil {
		return err
	}
	defer conn.Close()

	log.Printf("[http-server] Http server starting... [%s]", config.Address)
	return http.Serve(conn, corsMux)
}
