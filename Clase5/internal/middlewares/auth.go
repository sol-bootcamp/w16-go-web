package middlewares

import (
	"bootcamp-web/pkg/web"
	"net/http"
	"os"
)

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != os.Getenv("AUTH_SECRET") {
			web.Error(w, http.StatusUnauthorized, "Unauthorized")
			return
		}
		// Llamar al siguiente handler en la cadena
		next.ServeHTTP(w, r)
	})
}
