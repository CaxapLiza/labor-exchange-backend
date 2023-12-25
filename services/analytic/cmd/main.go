package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/student/labor-exchange-backend/services/analytic/internal/handler"
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

	router.HandleFunc("/analytics/vacancy", handler.GetVacancyResponsesCount).Methods("GET", "OPTIONS")
	router.HandleFunc("/analytics/employer", handler.GetEmployerResponsesCount).Methods("GET", "OPTIONS")
	router.HandleFunc("/analytics/job_seeker", handler.GetJobSeekerFavoritesCount).Methods("GET", "OPTIONS")

	log.Fatal(http.ListenAndServe(":8001", router))
}
