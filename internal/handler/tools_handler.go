package handler

import (
	"net/http"

	"devcortex.ai/internal/model"
	"devcortex.ai/internal/view"
)

// tools, sistemdeki tüm araçların listesidir.

func Tools(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Tüm Araçlar",
		FeaturedTools: model.FeaturedTools,
		OtherTools:    model.OtherTools,
	}
	view.Render(w, r, "tools.html", data)
}
