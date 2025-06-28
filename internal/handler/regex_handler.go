package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func RegexTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Regex Tester",
	}
	view.Render(w, r, "regex-tester.html", data)
}
