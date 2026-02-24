package main

import (
	"4-order-api/config"
	"4-order-api/internal/product"
	"4-order-api/pkg/db"
	"fmt"
	"net/http"
)

func main() {
	conf := config.LoadConfig()
	db := db.NewDb(conf.DSN)

	router := http.NewServeMux()

	productRepo := product.NewRepo(&product.ProductRepoDeps{Db: db})
	product.NewHandler(router, &product.HandlerDeps{Repo: productRepo})

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is running")
	server.ListenAndServe()
}
