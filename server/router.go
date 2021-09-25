package server

import (
	"github.com/julienschmidt/httprouter"
	"github.com/phthaocse/user-service-go/server/middleware"
)

func (srv *server) setUpRouter() {
	srv.router = httprouter.New()
	srv.router.POST("/user/register", middleware.HttprouterWrapper(middleware.Adapt(srv.Register(), middleware.LogToConsole(srv.Log))))
}
