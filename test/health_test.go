package test

import (
	"fmt"
	"testing"

	"github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	fmt.Println("Health check for localhost:8080")

	client := resty.New()
	resp, err := client.R().Get("http://localhost:8080/alive")
	if err != nil {
		t.Fail()
	}

	assert.Equal(t, 200, resp.StatusCode())
}
