package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func SQLAutopsyTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "SQL Query Autopsy",
		ToolSpecificData: make(map[string]interface{}),
	}

	data.ToolSpecificData.(map[string]interface{})["SQLQuery"] = "SELECT * FROM users WHERE id = 1;"
	data.ToolSpecificData.(map[string]interface{})["Analysis"] = "This query selects a specific user from the 'users' table. Ensure there is an index on the 'id' column for performance."

	view.Render(w, r, "sql-autopsy.html", data)
}
