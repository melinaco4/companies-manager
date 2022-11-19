package database

import (
	"context"
	"testing"

	"github.com/melinaco4/companies-manager/internal/company"
	"github.com/stretchr/testify/assert"
)

func TestCompanyDB(t *testing.T) {
	t.Run("test creation of company", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmpn, err := db.PostCompany(context.Background(), company.Company{
			Name:              "Google",
			Description:       "search engine",
			AmountofEmployees: 13245,
			Registered:        false,
			Type:              "Corporations",
		})
		assert.NoError(t, err)

		newCmpn, err := db.GetCompany(context.Background(), cmpn.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Google", newCmpn.Name)
	})

	t.Run("test deleting company", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmpn, err := db.PostCompany(context.Background(), company.Company{
			Name:              "Yahoo",
			Description:       "search engine",
			AmountofEmployees: 1234,
			Registered:        false,
			Type:              "Corporations",
		})
		assert.NoError(t, err)

		err = db.DeleteCompany(context.Background(), cmpn.ID)
		assert.NoError(t, err)

		_, err = db.GetCompany(context.Background(), cmpn.ID)
		assert.Error(t, err)
	})

	t.Run("test update comment", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmpn, err := db.PostCompany(context.Background(), company.Company{
			Name:              "Facebook",
			Description:       "social media",
			AmountofEmployees: 45678,
			Registered:        false,
			Type:              "Corporations",
		})
		assert.NoError(t, err)

		cmpn.Name = "Metaverse"
		cmpn, err = db.UpdateCompany(context.Background(), cmpn.ID, cmpn)
		assert.NoError(t, err)

		newCmt, err := db.GetCompany(context.Background(), cmpn.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Metaverse", newCmt.Name)
	})

	t.Run("test getting company", func(t *testing.T) {
		db, err := NewDatabase()
		assert.NoError(t, err)
		cmpn, err := db.PostCompany(context.Background(), company.Company{
			Name:              "Linkedin",
			Description:       "social media",
			AmountofEmployees: 3456,
			Registered:        false,
			Type:              "Corporations",
		})
		assert.NoError(t, err)

		newCmpn, err := db.GetCompany(context.Background(), cmpn.ID)
		assert.NoError(t, err)
		assert.Equal(t, "Linkedin", newCmpn.Name)

	})
}
