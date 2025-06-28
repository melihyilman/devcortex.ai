package handler

import (
	"net/http"

	"devcortex.ai/internal/model"

	"devcortex.ai/internal/view"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Ana Sayfa",
		Tools: model.Tools,
	}
	view.Render(w, r, "home.html", data)
}
