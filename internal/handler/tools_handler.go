package handler

import (
	"net/http"

	"devcortex.ai/internal/model"
	"devcortex.ai/internal/view"
)

func Tools(w http.ResponseWriter, r *http.Request) {
	
	categorizedTools := make(map[string][]model.Tool)
	for _, tool := range model.OtherTools {
		categorizedTools[tool.Category] = append(categorizedTools[tool.Category], tool)
	}

	data := &view.PageData{
		Title:         "All Tools",
		FeaturedTools: model.FeaturedTools,
		
		ToolSpecificData: categorizedTools,
	}
	view.Render(w, r, "tools.html", data)
}
