package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

func LoggingMiddleware(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		fmt.Printf("Iniciando ejecución de la petición %s %s, %v \n", r.Method, r.URL.Path, start)

		// Llamar al siguiente handler en la cadena
		handler.ServeHTTP(w, r)

		fmt.Printf("Finalizando ejecución de la petición %s %s, %v \n", r.Method, r.URL.Path, time.Since(start))

	})
}
