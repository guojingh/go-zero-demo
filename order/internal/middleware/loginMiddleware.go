package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type LoginMiddleware struct {
}

func NewLoginMiddleware() *LoginMiddleware {
	return &LoginMiddleware{}
}

func (m *LoginMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// TODO generate middleware implement function, delete after code implementation

		httpx.OkJson(w, "请先登陆")
		//return

		// Passthrough to next handler if need
		next(w, r)
	}
}
