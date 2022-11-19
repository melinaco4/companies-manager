package http

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"net/http"

	"github.com/gorilla/mux"
	"github.com/melinaco4/companies-manager/internal/company"
)

type Handler struct {
	Router  *mux.Router
	Service CompanyService
	Server  *http.Server
}

type Response struct {
	Message string
	Error   string
}

func NewHandler(service *company.Service) *Handler {
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

func (h *Handler) Serve() error {
	go func() {
		if err := h.Server.ListenAndServe(); err != nil {
			log.Println(err.Error())
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shut down gracefully")
	return nil
}

func (h *Handler) mapRoutes() {

	//h.Router = mux.NewRouter()
	h.Router.HandleFunc("/api/company/{id}", h.GetCompany).Methods("GET")
	h.Router.HandleFunc("/api/company", h.PostCompany).Methods("POST")
	h.Router.HandleFunc("/api/company/{id}", h.UpdateCompany).Methods("PATCH")
	h.Router.HandleFunc("/api/company/{id}", h.DeleteCompany).Methods("DELETE")
}
