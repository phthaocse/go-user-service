package handlers

import (
	"context"
	"fmt"
	"github.com/phthaocse/user-service-go/models"
	"net/http"
)

func (handler *Handler) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler.Log.Println("Register")

		count, err := models.Users().Count(context.Background(), handler.Db)
		if err != nil {
			fmt.Println(err)
			return
		}
		handler.Log.Println("count: ", count)
		w.WriteHeader(http.StatusOK)
	}
}
