package handler

import (
	"net/http"

	"devcortex.ai/internal/model"
	"devcortex.ai/internal/view"
)

func Tools(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "All Tools",
		FeaturedTools: model.FeaturedTools,
		OtherTools:    model.OtherTools,
	}
	view.Render(w, r, "tools.html", data)
}