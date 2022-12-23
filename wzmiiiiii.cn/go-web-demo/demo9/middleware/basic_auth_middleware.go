// @author: wzmiiiiii
// @since: 2022/12/23 19:42:07
// @desc: TODO

package middleware

import (
	"net/http"
)

type BasicAuthMiddleware struct {
	Next http.Handler
}

func (m *BasicAuthMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if m.Next == nil {
		m.Next = http.DefaultServeMux
	}
	//if r.Method != http.MethodGet {
	//	username, passwd, ok := r.BasicAuth()
	//	if !ok {
	//		log.Println("Error parsing basic auth")
	//		w.WriteHeader(http.StatusUnauthorized)
	//		return
	//	}
	//	if username != "admin" {
	//		log.Println("Username is not correct")
	//		w.WriteHeader(http.StatusUnauthorized)
	//		return
	//	}
	//	if passwd != "123456" {
	//		log.Println("Password is not correct")
	//		w.WriteHeader(http.StatusUnauthorized)
	//		return
	//	}
	//}
	m.Next.ServeHTTP(w, r)
}
