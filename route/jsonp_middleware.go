package route

import (
	"net/http"
)

// JsonpMiddleware wraps the json output with callback function name
func JsonpMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	cb := r.URL.Query().Get("callback")
	if cb != "" {
		w.Write([]byte(cb))
		w.Write([]byte("("))
	}

	next(w, r)

	if cb != "" {
		w.Write([]byte(")"))
	}
}
