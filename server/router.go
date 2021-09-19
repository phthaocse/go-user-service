package server

import (
	"github.com/julienschmidt/httprouter"
	"github.com/phthaocse/user-service-go/server/middleware"
)

func (srv *Server) setUpRouter() {
	srv.router = httprouter.New()
	srv.router.POST("/user/register", middleware.HttprouterWrapper(middleware.Adapt(srv.handler.Register(), middleware.LogToConsole(srv.handler.Log))))
}
