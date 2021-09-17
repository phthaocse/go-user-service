package handlers

import (
	"context"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/phthaocse/user-service-go/models"
	"net/http"
)

func (handler *Handler) Register(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	fmt.Println("Register")

	count, err := models.Users().Count(context.Background(), handler.Db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("count: ", count)
	w.WriteHeader(http.StatusOK)
}