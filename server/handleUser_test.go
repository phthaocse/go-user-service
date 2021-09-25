package server

import (
	"github.com/joho/godotenv"
	"github.com/matryer/is"
	"github.com/phthaocse/user-service-go/config"
	"github.com/phthaocse/user-service-go/db"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegister(t *testing.T) {
	is := is.New(t)
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	srv := CreateServer()
	dbCon, dbTeardown, err := db.SetUp(config.GetTestSrvConfig(), srv.Log)
	defer dbTeardown()
	if err != nil {
		srv.Log.Fatal(err)
	}
	srv.db = dbCon

	r := httptest.NewRequest("POST", "/user/register", nil)
	w := httptest.NewRecorder()
	srv.ServeHTTP(w, r)
	is.Equal(w.Code, http.StatusBadRequest)
}
