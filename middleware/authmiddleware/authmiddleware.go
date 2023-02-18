package authmiddleware

import (
	"log"
	"net/http"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		r.Header.Add("key", "value")

		next.ServeHTTP(w, r)

		log.Printf(
			"%s %s %s %s",
			r.Method,
			r.RequestURI,
			time.Since(startTime),
			r.RemoteAddr,
		)
	})
}
