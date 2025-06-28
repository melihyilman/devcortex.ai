package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func MinifierTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "HTML/CSS/JS Minifier",
	}
	view.Render(w, r, "minifier.html", data)
}
