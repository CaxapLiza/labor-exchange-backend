package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/student/labor-exchange-backend/services/account/internal/handler"
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

	router.HandleFunc("/accounts", handler.GetList).Methods("GET", "OPTIONS")
	router.HandleFunc("/accounts/{id}", handler.Get).Methods("GET", "OPTIONS")
	router.HandleFunc("/accounts", handler.Create).Methods("POST", "OPTIONS")
	router.HandleFunc("/accounts/{id}", handler.Update).Methods("PUT", "OPTIONS")
	router.HandleFunc("/accounts/{id}", handler.Delete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/accounts", handler.Authenticate).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8080", router))
}
