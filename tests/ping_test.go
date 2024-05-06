package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mohamedhani/test-github-action/router"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	var dat map[string]interface{}
	if err := json.Unmarshal(w.Body.Bytes(), &dat); err != nil {
		panic(err)
	}

	assert.Equal(t, "pong", dat["message"])
}
