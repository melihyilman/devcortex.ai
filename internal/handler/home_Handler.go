package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Home Page",
	}
	view.Render(w, r, "home.html", data)
}