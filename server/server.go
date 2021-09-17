package server

import (
	"database/sql"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/phthaocse/user-service-go/config"
	"github.com/phthaocse/user-service-go/db"
	"github.com/phthaocse/user-service-go/handlers"
	"net/http"
)

type Server struct {
	router *httprouter.Router
	handler *handlers.Handler
}

func createServer() *Server {
	s := &Server{}
	s.handler = handlers.CreateHandler()
	s.setUpRouter()
	return s
}

func (srv *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	srv.router.ServeHTTP(rw, req)
}

func Start() {
	var dbConnection *sql.DB
	globalConfig, err := config.SrvConfigSetup()
	if err != nil {
		return
	}
	dbConnection, err = db.SetUp(globalConfig)
	srv := createServer()
	srv.handler.Db = dbConnection

	fmt.Println("Start running ...", globalConfig.ServerPort)
	err = http.ListenAndServe(globalConfig.ServerPort, srv)
	if err != nil {
		return
	}
}