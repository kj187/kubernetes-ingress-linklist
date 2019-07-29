package links

import (
	"io"
	"net/http"
)

// LinksHandler ...
func LinksHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, `Linklist`)
}
