package middleware

import (
	"light-cloud/src/core/helper"
	"net/http"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if auth == "" {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte("Unauthorized"))
			if err != nil {
				return
			}
			return
		}
		userClaim, err := helper.AnalyzeToken(auth)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			_, err := w.Write([]byte(err.Error()))
			if err != nil {
				return
			}
			return
		}
		// 一切成功，将基础信息写入请求头，便于调取
		r.Header.Set("UserId", string(rune(userClaim.Id)))
		r.Header.Set("UserIdentity", userClaim.Identity)
		r.Header.Set("UserName", userClaim.Name)
		next(w, r)
	}
}
