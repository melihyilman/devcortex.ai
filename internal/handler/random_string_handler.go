package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func RandomStringTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Random String Generator",
	}
	view.Render(w, r, "random-string.html", data)
}
