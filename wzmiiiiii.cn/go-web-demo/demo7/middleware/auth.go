// @author: wzmiiiiii
// @since: 2022/12/23 02:13:45
// @desc: TODO

package middleware

import "net/http"

type AuthMiddleware struct {
	Next http.Handler
}

func (am *AuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if am.Next == nil {
		am.Next = http.DefaultServeMux
	}
	auth := r.Header.Get("Authorization")
	if auth != "" {
		am.Next.ServeHTTP(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
