package server

import (
	"database/sql"
	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/phthaocse/user-service-go/config"
	"github.com/phthaocse/user-service-go/db"
	"log"
	"net/http"
	"os"
)

type server struct {
	router *httprouter.Router
	db     *sql.DB
	Log    *log.Logger
}

func CreateServer() *server {
	srv := &server{}
	srv.Log = log.New(os.Stdout, "", log.LstdFlags)
	srv.setUpRouter()
	return srv
}

func (srv *server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	srv.router.ServeHTTP(rw, req)
}

func Start() {

	if utils.GetEnv("ENV", "") == "DEV" {
		err := godotenv.Load()
		if err != nil {
			return
		}
	}

	var dbConnection *sql.DB
	globalConfig := config.GetSrvConfig()

	srv := CreateServer()

	dbConnection, dbTeardown, err := db.SetUp(globalConfig, srv.Log)
	defer dbTeardown()
	srv.db = dbConnection

	srv.Log.Println("\033[46;1m", "Start running server on port", globalConfig.ServerPort, "\033[0m")
	err = http.ListenAndServe(globalConfig.ServerPort, srv)
	if err != nil {
		srv.Log.Println(err)
		return
	}
}
