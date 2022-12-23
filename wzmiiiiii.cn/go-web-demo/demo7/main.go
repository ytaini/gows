// @author: wzmiiiiii
// @since: 2022/12/23 02:08:01
// @desc: //TODO

package main

import (
	"encoding/json"
	"net/http"
	"time"

	"wzmiiiiii.cn/gwd/demo7/middleware"
)

func main() {

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		c := Company{
			ID:      123,
			Name:    "Google",
			Country: "USA",
		}
		time.Sleep(4 * time.Second)
		enc := json.NewEncoder(w)
		enc.Encode(c)
	})

	http.ListenAndServe(":8080", &middleware.TimeoutMiddleware{
		Next: new(middleware.AuthMiddleware),
	})
}
