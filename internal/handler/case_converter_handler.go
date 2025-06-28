package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func CaseConverterTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Case Converter",
	}
	view.Render(w, r, "case-converter.html", data)
}
