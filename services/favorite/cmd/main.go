package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/student/labor-exchange-backend/services/favorite/internal/handler"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()

	corsMiddleware := handlers.CORS(
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
	)

	router.Use(corsMiddleware)

	router.HandleFunc("/favorites/{id}", handler.GetList).Methods("GET", "OPTIONS")
	router.HandleFunc("/favorites", handler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/favorites/{id}", handler.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/favorites/{id}", handler.Delete).Methods("DELETE", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8083", router))
}
