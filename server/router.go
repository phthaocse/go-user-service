package server

import (
	"github.com/julienschmidt/httprouter"
)

func (srv *Server) setUpRouter() {
	srv.router = httprouter.New()
	srv.router.GET("/user/register", srv.handler.Register)
}
