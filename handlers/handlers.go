package handlers

import (
	"encoding/json"
	"net/http"
	"program/model"
	"program/users"

	"github.com/gorilla/mux"
)

type apiHandler struct {
	UserServer *users.UserServer
}

func RetHandler(userServer *users.UserServer) *apiHandler {
	return &apiHandler{
		UserServer: userServer,
	}
}
func HandleRequest(h *apiHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/health/live", HealthCheck).Methods(http.MethodGet)
	r.HandleFunc("/user/create", h.CreateUser).Methods(http.MethodPost)
	return r
}
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(http.StatusOK)
}
func (h *apiHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = model.User.ValidateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = h.UserServer.CreateUser(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusCreated)
}
