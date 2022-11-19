package test

import (
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestPostCompany(t *testing.T) {
	client := resty.New()
	resp, err := client.R().
		SetBody(`{
		"name":"TestCompany",
		"description":"",
		"amountofemployees":1112,
		"registered":true,
		"type": "NonProfit"}`).
		Post("http://localhost:8080/api/company")
	assert.NoError(t, err)

	assert.Equal(t, 200, resp.StatusCode())
}
