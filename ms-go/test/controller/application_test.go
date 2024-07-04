package controller

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"ms-go/router"

	"github.com/gin-gonic/gin"
	"github.com/likexian/gokit/assert"
)

func TestIndexHome(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"message": "[ms-go] | Success",
		"status":  200,
	}

	// Grab our router
	router := router.SetupRouter()

	// Perform a GET request with that handler.
	req, _ := http.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	// Assert we encoded correctly,
	assert.Equal(t, http.StatusOK, w.Code)

	// Convert the correct object response to a map
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON response: %v", err)
	}

	// Make some assertions on the correctness of the response.
	assert.Equal(t, body["message"], response["message"])
}
