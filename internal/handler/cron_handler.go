package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func CronTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Zaman Gezgini (Cron)",
		ToolSpecificData: nil,
	}
	view.Render(w, r, "cron.html", data)
}
