package main

import (
	"3-validation-api/config"
	"3-validation-api/internal/verify"
	"fmt"
	"net/http"
)

func main() {
	config := config.LoadConfig()

	router := http.NewServeMux()

	verify.New(router, &verify.VerifyHandlerDeps{
		Config: config,
	})

	server := &http.Server{
		Addr:    ":8081",
		Handler: router,
	}

	fmt.Println("Server is running")
	server.ListenAndServe()
}
