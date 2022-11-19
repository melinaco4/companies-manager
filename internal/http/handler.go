package http

import (
	"encoding/json"

	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/melinaco4/companies-manager/internal/company"
)

type Handler struct {
	Router  *mux.Router
	Service *company.Service
}

type Response struct {
	Message string
	Error   string
}

func NewHandler(service *company.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) mapRoutes() {

	h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/company/{id}", h.GetCompany).Methods("GET")
	h.Router.HandleFunc("/api/company", h.PostCompany).Methods("POST")
	h.Router.HandleFunc("/api/company{id}", h.UpdateCompany).Methods("PATCH")
	h.Router.HandleFunc("/api/company{id}", h.DeleteCompany).Methods("DELETE")
}

func (h *Handler) PostCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	var company company.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
	}

	company, err := h.Service.CreateCompany(company)
	if err != nil {
		sendErrorResponse(w, "Failed to create new company", err)
	}
	if err := json.NewEncoder(w).Encode(company); err != nil {
		panic(err)
	}
}

func (h *Handler) GetCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]

	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}
	company, err := h.Service.GetCompany(uint(i))
	if err != nil {
		sendErrorResponse(w, "Failed Retrieving Company By ID", err)
	}

	if err := json.NewEncoder(w).Encode(company); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	id := vars["id"]
	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse uint from ID", err)
	}

	var company company.Company
	if err := json.NewDecoder(r.Body).Decode(&company); err != nil {
		sendErrorResponse(w, "Failed to decode JSON Body", err)
	}

	company, err = h.Service.UpdateCompany(uint(companyID), company)
	if err != nil {
		sendErrorResponse(w, "Failed to update company", err)
	}
	if err := json.NewEncoder(w).Encode(company); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	vars := mux.Vars(r)
	id := vars["id"]
	companyID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Failed to parse uint from ID", err)
	}

	err = h.Service.DeleteCompany(uint(companyID))
	if err != nil {
		sendErrorResponse(w, "Failed to delete comment by company ID", err)
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Company Deleted"}); err != nil {
		panic(err)
	}

}

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}
