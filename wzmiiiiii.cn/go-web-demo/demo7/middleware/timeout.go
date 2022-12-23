// @author: wzmiiiiii
// @since: 2022/12/23 02:40:37
// @desc: TODO

package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm *TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}

	ctx, cancel := context.WithTimeout(r.Context(), 3*time.Second)
	defer cancel()

	r = r.WithContext(ctx)

	ch := make(chan struct{})

	go func() {
		tm.Next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()

	select {
	case <-ch:
		return
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}
}
