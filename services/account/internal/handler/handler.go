package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/student/labor-exchange-backend/services/account/internal"
	"github.com/student/labor-exchange-backend/services/account/internal/repository"
	"github.com/student/labor-exchange-backend/services/common"
	"log"
	"net/http"
	"strconv"
)

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func GetList(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	accounts, err := repo.GetList()
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		log.Println("Get Error:", err)
		return
	}

	response, err := json.Marshal(accounts)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Get(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid ID:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	accounts, err := repo.Get(id)
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		log.Println("Get Error:", err)
		return
	}

	response, err := json.Marshal(accounts)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}

func Create(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var requestBody struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Invalid request body:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	account := &internal.Account{Login: requestBody.Login, Password: requestBody.Password,
		Role: requestBody.Role}

	if err := repo.Create(account); err != nil {
		http.Error(w, "Error creating", http.StatusInternalServerError)
		log.Println("Create Error:", err)
		return
	}

	response, err := json.Marshal(account)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

func Update(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid ID:", err)
		return
	}

	var requestBody struct {
		Login    string `json:"login"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Invalid request body:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	if err := repo.Update(id, requestBody.Login, requestBody.Password,
		requestBody.Role); err != nil {
		http.Error(w, "Error updating", http.StatusInternalServerError)
		log.Println("Update Error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		log.Println("Invalid ID:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	if err := repo.Delete(id); err != nil {
		http.Error(w, "Error deleting", http.StatusInternalServerError)
		log.Println("Delete Error:", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func Authenticate(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	var requestBody struct {
		Login    string `json:"login"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Println("Invalid request body:", err)
		return
	}

	db, err := common.NewDatabase()
	if err != nil {
		http.Error(w, "Unable to connect to the database", http.StatusInternalServerError)
		log.Println("Connection Error:", err)
		return
	}
	defer db.Close()

	repo := repository.NewRepository(db)

	account, err := repo.Authenticate(requestBody.Login, requestBody.Password)
	if err != nil {
		http.Error(w, "Authentication failed", http.StatusUnauthorized)
		log.Println("Authentication Error:", err)
		return
	}

	response := struct {
		ID   int    `json:"id"`
		Role string `json:"role"`
	}{
		ID:   account.ID,
		Role: account.Role,
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error encoding JSON", http.StatusInternalServerError)
		log.Println("JSON Error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
