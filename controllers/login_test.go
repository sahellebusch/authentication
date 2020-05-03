package controllers_test

import (
	"authentication/controllers"
	"authentication/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"authentication/models"
	// "github.com/jinzhu/gorm"
	"bytes"
	"encoding/json"
	. "github.com/smartystreets/goconvey/convey"
)

// TestStatusRoute
func TestLoginRoute(t *testing.T) {
	Convey("With a malformed request", t, func() {
		db := models.SetupModels()
		// testUser := models.User{Username: "jimcarey", Password: "legend"}
		// db.Create(&testUser)

		defer db.Close()
		router := router.SetupRouter(db)
		Convey("With no username", func() {
			withoutUsername, _ := json.Marshal(&controllers.Credentials{Password: "test"})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(withoutUsername))
			router.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, http.StatusBadRequest)
			So(w.Body.String(), ShouldContainSubstring, "Error:Field validation for 'Username'")

		})

		Convey("With no password", func() {
			withoutUsername, _ := json.Marshal(&controllers.Credentials{Username: "test"})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(withoutUsername))
			router.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, http.StatusBadRequest)
			So(w.Body.String(), ShouldContainSubstring, "Error:Field validation for 'Password'")

		})
	})
}
