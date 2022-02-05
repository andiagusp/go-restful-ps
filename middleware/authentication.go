package middleware

import (
	"golang-restful/helper"
	"golang-restful/model/web"
	"net/http"
)

type AuthMiddleware struct {
	Handler http.Handler
}

func NewAuthMiddleware(handler http.Handler) *AuthMiddleware {
	return &AuthMiddleware{
		Handler: handler,
	}
}

func (middleware *AuthMiddleware) ServeHTTP(wr http.ResponseWriter, req *http.Request) {
	key := req.Header.Get("X-API-Key")

	if key == "secret" {
		middleware.Handler.ServeHTTP(wr, req)
	} else {
		wr.Header().Set("Content-Type", "application/json")
		wr.WriteHeader(http.StatusUnauthorized)
		webResponse := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "Unauthorized",
		}
		helper.WriteEncodeJson(wr, webResponse)
	}

}
