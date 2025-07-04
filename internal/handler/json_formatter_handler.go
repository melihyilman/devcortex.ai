package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func JsonFormatter(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "JSON Formatter",
	}
	view.Render(w, r, "json-formatter.html", data)
}