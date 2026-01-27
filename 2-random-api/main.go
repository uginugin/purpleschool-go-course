package main

import (
	"math/rand/v2"
	"net/http"
	"strconv"
)

func main() {
	router := http.NewServeMux()
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	router.HandleFunc("/random", handleRandom)

	server.ListenAndServe()

}

func handleRandom(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(strconv.Itoa(rand.IntN(7))))
}
