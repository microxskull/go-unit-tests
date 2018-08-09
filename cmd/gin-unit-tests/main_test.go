package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPingGETRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, `{"message":"pong"}`, w.Body.String())
}

func TestPingPOSTRoute(t *testing.T) {
	router := setupRouter()
	w := httptest.NewRecorder()
	message := time.Now().Format("2006-01-02 15:04:05")
	req, _ := http.NewRequest("POST", "/ping",
		strings.NewReader(fmt.Sprintf(`{"message":"%v"}`, message)))
	router.ServeHTTP(w, req)
	// Check test
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintf(`{"message":"pong %v"}`, message), w.Body.String())
}
