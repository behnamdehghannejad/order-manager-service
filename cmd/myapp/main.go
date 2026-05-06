package main

import (
	"log"
	"net/http"
	"order-management-service/internal/config"
	"order-management-service/internal/handler"
	"order-management-service/internal/infra/repository"
	"order-management-service/internal/service"

	"gorm.io/gorm"
)

func main() {
	db := startService()

	orderRepository := repository.NewOrderRepository(db)
	orderService := service.NewOrderService(orderRepository)
	orderHandler := handler.NewOrderHandler(orderService)

	mux := http.NewServeMux()

	handelRequests(mux, orderHandler)

	http.ListenAndServe(":8080", mux)
}

func handelRequests(mux *http.ServeMux, orderHandler *handler.OrderHandler) {
	mux.HandleFunc("POST /oms/api/orders", orderHandler.CreateOrder)
	mux.HandleFunc("GET /oms/api/orders/{id}", orderHandler.GetOrder)
}

func startService() *gorm.DB {
	cfg := loadConfiguration()
	return initPostgres(cfg)
}

func loadConfiguration() *config.Config {
	cfg, err := config.LoadConfig("config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}

func initPostgres(cfg *config.Config) *gorm.DB {
	db, err := config.NewPostgres(cfg)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&repository.OrderEntity{})

	return db
}
