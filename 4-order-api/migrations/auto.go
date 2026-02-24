package main

import (
	"4-order-api/config"
	"4-order-api/internal/product"
	"4-order-api/pkg/db"
)

func main() {
	dsn := config.LoadConfig().DSN

	db := db.NewDb(dsn)

	db.AutoMigrate(&product.Product{})
}
