package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func CronExplorerTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "Zaman Gezgini (Cron)",
		ToolSpecificData: make(map[string]interface{}),
	}

	// Placeholder for cron logic
	// In a real application, you would parse the cron expression
	// and extract information like next run times, etc.
	data.ToolSpecificData.(map[string]interface{})["CronExpression"] = "* * * * *"
	data.ToolSpecificData.(map[string]interface{})["Description"] = "Her dakika çalışır."

	view.Render(w, r, "cron-explorer.html", data)
}
