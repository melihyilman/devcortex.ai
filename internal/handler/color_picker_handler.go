package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func ColorPickerTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Color Picker / Converter",
	}
	view.Render(w, r, "color-picker.html", data)
}
