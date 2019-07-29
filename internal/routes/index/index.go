package index

import (
	"net/http"
)

// IndexHandler ...
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "//"+r.Host+"/default", http.StatusPermanentRedirect)
}
