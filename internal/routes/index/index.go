package index

import (
	"io"
	"net/http"
)

// IndexHandler ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `Index`)
}
