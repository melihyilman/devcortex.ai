package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func ExtensionsHandler(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Extensions",
	}

	view.Render(w, r, "extensions.html", data)
}