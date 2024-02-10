package delivery

import (
	"database/sql"
	"fmt"
	"main/config"
	"main/delivery/controller"
	"main/repository"
	"main/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	productUc usecase.ProductUsecase
	engine    *gin.Engine
	port      string
}

func (s *Server) initRouter() {
	rg := s.engine.Group(config.APiGroup)
	controller.NewProductController(s.productUc, rg).Route()
}

func (s *Server) Run() {
	s.initRouter()
	if err := s.engine.Run(s.port); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.port, err.Error()))
	}
}

func NewSever() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbName=%s sslmode=disable", cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic("connection error")
	}

	//repo
	productRepo := repository.NewProductRepository(db)

	//usecase
	productUc := usecase.NewProductUsecase(productRepo)

	engine := gin.Default()
	port := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{productUc: productUc, engine: engine, port: port}

}
