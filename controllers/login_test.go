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
	db := models.SetupModels()
	defer db.Close()
	router := router.SetupRouter(db)
	db.Create(&models.User{Password: "super", Username: "tomcollins"})

	Convey("Creates a new JWT token", t, func() {
		creds, _ := json.Marshal(&controllers.Credentials{Username: "tomcollins", Password: "super"})
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(creds))
		router.ServeHTTP(w, req)

		So(w.Code, ShouldEqual, http.StatusOK)
		So(w.Header().Get("Content-Type"), ShouldEqual, "application/json; charset=utf-8")
		var respJSON controllers.JWTResponse
		json.Unmarshal([]byte(w.Body.String()), &respJSON)
		So(respJSON.JWT, ShouldNotBeEmpty)
	})

	Convey("With a malformed request", t, func() {
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

	Convey("With invalid username and password combination", t, func() {
		Convey("invalid username", func() {
			wrongUsername, _ := json.Marshal(&controllers.Credentials{Username: "wrongo", Password: "super"})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(wrongUsername))
			router.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, http.StatusForbidden)
			So(w.Body.String(), ShouldContainSubstring, "Invalid username and password combination")
		})

		Convey("invalid password", func() {
			wrongPass, _ := json.Marshal(&controllers.Credentials{Username: "tomcollins", Password: "nope"})
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/v1/login", bytes.NewBuffer(wrongPass))
			router.ServeHTTP(w, req)

			So(w.Code, ShouldEqual, http.StatusForbidden)
			So(w.Body.String(), ShouldContainSubstring, "Invalid username and password combination")
		})
	})
}
