package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func MarkdownPreviewerTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Markdown Previewer",
	}
	view.Render(w, r, "markdown-previewer.html", data)
}
