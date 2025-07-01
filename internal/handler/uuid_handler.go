package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
	"github.com/google/uuid"
)

func UUIDTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "UUID Generator",
		ToolSpecificData: make(map[string]interface{}),
	}

	if r.Method == http.MethodPost {
		uuid := uuid.New().String()
		data.ToolSpecificData.(map[string]interface{})["UUID"] = uuid
	}

	view.Render(w, r, "uuid-generator.html", data)
}
