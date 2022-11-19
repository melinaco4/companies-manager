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

	// Sets up our middleware functions
	h.Router.Use(JSONMiddleware)
	// we also want to log every incoming request
	h.Router.Use(LoggingMiddleware)
	// We want to timeout all requests that take longer than 15 seconds
	h.Router.Use(TimeoutMiddleware)

	h.Server = &http.Server{
		Addr:         "0.0.0.0:8080",
		Handler:      h.Router,
		WriteTimeout: time.Second * 10,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Second * 60,
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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	h.Server.Shutdown(ctx)

	log.Println("shut down gracefully")
	return nil
}

func (h *Handler) mapRoutes() {

	h.Router.HandleFunc("/alive", h.Health).Methods("GET")
	h.Router.HandleFunc("/api/company/{id}", JWTAuth(h.GetCompany)).Methods("GET")
	h.Router.HandleFunc("/api/company", JWTAuth(h.PostCompany)).Methods("POST")
	h.Router.HandleFunc("/api/company/{id}", JWTAuth(h.UpdateCompany)).Methods("PATCH")
	h.Router.HandleFunc("/api/company/{id}", JWTAuth(h.DeleteCompany)).Methods("DELETE")
}
