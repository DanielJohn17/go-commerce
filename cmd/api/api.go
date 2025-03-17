package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/DanielJohn17/go-commerce/cmd/api/service/cart"
	"github.com/DanielJohn17/go-commerce/cmd/api/service/order"
	"github.com/DanielJohn17/go-commerce/cmd/api/service/product"
	"github.com/DanielJohn17/go-commerce/cmd/api/service/user"
	"github.com/gin-gonic/gin"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := gin.Default()
	subRouter := router.Group("/api/v1")

	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subRouter)

	productStore := product.NewStore(s.db)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subRouter)

	orderStore := order.NewStore(s.db)

	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subRouter)

	log.Println("Listening on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
