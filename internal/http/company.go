package http

import (
	"context"
	"encoding/json"
	"net/http"

	"log"

	"github.com/melinaco4/companies-manager/internal/company"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type CompanyService interface {
	PostCompany(context.Context, company.Company) (company.Company, error)
	GetCompany(ctx context.Context, ID string) (company.Company, error)
	UpdateCompany(ctx context.Context, ID string, newCmpn company.Company (company.Company, error)
	DeleteCompany(ctx context.Context, ID string) error
}

type Response struct {
	Message string
}

type PostCompanyRequest struct {
	Name  string `json:"slug" validate:"required"`
	Description string `json:"author"`
	AmountofEmployees   int `json:"amountofemployees" validate:"required"`
	Registered  string `json:"registered" validate:"required"`
	Type   string `json:"body" validate:"required"`
}

func (convertPostCompanyRequestToCompanyc PostCompanyRequest) company.Company{
	return company.Company{
		Name:   c.Name,
		Description: c.Description,
		AmountofEmployees:   c.AmountofEmployees,
		Registered: c.Registered,
		Type: c.Type,
	}
}

func (h *Handler) PostCompany(w http.ResponseWriter, r *http.Request) {
	var cmpn PostCompanyRequest
	if err := json.NewDecoder(r.Body).Decode(&cmpn); err != nil {
		return
	}

	validate := validator.New()
	err := validate.Struct(cmpn)
	if err != nil {
		http.Error(w, "not a valid company", http.StatusBadRequest)
		return
	}

	convertedCompany:= convertPostCompanyRequestToCompany(cmpn)
	postedCompany, err := h.Service.PostCompany(r.Context(), convertedCompany)
	if err != nil {
		log.Print(err)
		return
	}

	if err := json.NewEncoder(w).Encode(postedCompany); err != nil {
		panic(err)
	}

}

func (h *Handler) GetCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmpn, err := h.Service.GetCompany(r.Context(), id)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmpn company.Company
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmpn, err := h.Service.UpdateCompany(r.Context(), id, cmpn)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmpn); err != nil {
		panic(err)
	}

}

func (h *Handler) DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyID := vars["id"]

	if companyID == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteCompany(r.Context(), companyID)
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully deleted"}); err != nil {
		panic(err)
	}

}