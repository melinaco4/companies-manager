package http

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/melinaco4/companies-manager/internal/company"
	log "github.com/sirupsen/logrus"
)

type CompanyService interface{}

type Handler struct {
	Router  *mux.Router
	Service CompanyService
	Server  *http.Server
}

type Response struct {
	Message string
	Error   string
}

func NewHandler(service CompanyService) *Handler {
	h := &Handler{
		Service: service,
	}
	h.Router = mux.NewRouter()
	h.mapRoutes()

	h.Server = &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: h.Router,
	}

	return h
}

func (h *Handler) mapRoutes() {
	h.Router.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello world of things")
	})

	h.Router.HandleFunc("/api/company/{id}", h.GetCompany).Methods("GET")
	h.Router.HandleFunc("/api/company", h.PostCompany).Methods("POST")
	h.Router.HandleFunc("/api/company{id}", h.UpdateCompany).Methods("PATCH")
	h.Router.HandleFunc("/api/company{id}", h.DeleteCompany).Methods("DELETE")
}

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shutdown gracefully")

	return nil
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
