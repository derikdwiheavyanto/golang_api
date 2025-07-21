package middleware

import (
	"golang_restful_api/helper"
	"golang_restful_api/model/web"
	"net/http"
)

type AuthMiddleware struct {
	handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{handler: handler}
}

func (middleware *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("X-Api-Key") == "RAHASIA" {
		middleware.handler.ServeHTTP(w, r)
	} else {
		w.Header().Set("Content-Type", "aplication/json")
		w.WriteHeader(http.StatusUnauthorized)

		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}

		helper.WriteToResponseBody(w, webResponse)
	}
}
