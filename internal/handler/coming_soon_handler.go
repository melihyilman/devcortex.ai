package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func ComingSoon(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Coming Soon",
	}
	view.Render(w, r, "coming-soon.html", data)
}
