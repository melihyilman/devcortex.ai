package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func DiffCheckerTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Diff Checker",
	}
	view.Render(w, r, "diff-checker.html", data)
}
