package test

import (
	"authentication/router"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatusRoute(t *testing.T) {

	db, err := gorm.Open("sqlite3", "./test.db")
	if err != nil {
		panic("Failed to connect to database!")
	}

	defer db.Close()

	router := router.SetupRouter(db)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/status", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"status\":\"we good\"}", w.Body.String())
}
