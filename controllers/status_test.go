package controllers_test

import (
	"authentication/router"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"
	. "github.com/smartystreets/goconvey/convey"
)

func TestStatusRoute(t *testing.T) {
	Convey("When the server is running", t, func() {
		db, err := gorm.Open("sqlite3", "./test.db")
		if err != nil {
			panic("Failed to connect to database!")
		}

		defer db.Close()

		router := router.SetupRouter(db)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/status", nil)
		router.ServeHTTP(w, req)

		So(w.Code, ShouldEqual, 200)
		So(w.Body.String(), ShouldEqual, "{\"status\":\"we good\"}")
	})
}
