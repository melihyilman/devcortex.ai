package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func UUIDTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "UUID Generator",
	}
	view.Render(w, r, "uuid-generator.html", data)
}
