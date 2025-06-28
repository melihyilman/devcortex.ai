package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func URLEncoderTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "URL Encoder/Decoder",
	}
	view.Render(w, r, "url-encoder.html", data)
}
