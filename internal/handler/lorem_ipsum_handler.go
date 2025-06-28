package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func LoremIpsumTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Lorem Ipsum Generator",
	}
	view.Render(w, r, "lorem-ipsum.html", data)
}
