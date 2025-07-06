package handler

import (
	"net/http"
	"net/url"
)

// redirectToPageWithData takes the response writer, request, and a map of data.
// It constructs a redirect URL with the data as query parameters and performs a redirect.
func redirectToPageWithData(w http.ResponseWriter, r *http.Request, data map[string]string) {
	redirectURL, _ := url.Parse(r.URL.Path)
	query := redirectURL.Query()

	for key, value := range data {
		query.Set(key, value)
	}
	redirectURL.RawQuery = query.Encode()

	http.Redirect(w, r, redirectURL.String(), http.StatusSeeOther)
}
