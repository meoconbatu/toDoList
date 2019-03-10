package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	gintemplate "github.com/foolin/gin-template"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}

func getRouter(withTemplate bool) *gin.Engine {
	r := gin.Default()

	if withTemplate {
		r.HTMLRender = gintemplate.Default()
	}
	return r
}

func initDB() *gorm.DB {
	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	return db
}
func testHTTPResponse(t *testing.T, r *gin.Engine, req *http.Request, f func(w *httptest.ResponseRecorder) bool) {
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if !f(w) {
		t.Fail()
	}
}
