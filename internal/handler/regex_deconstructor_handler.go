package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func RegexDeconstructorTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title:            "Regex Deconstructor",
		ToolSpecificData: make(map[string]interface{}),
	}

	data.ToolSpecificData.(map[string]interface{})["RegexPattern"] = `^([a-zA-Z0-9._%+-]+)@([a-zA-Z0-9.-]+\.[a-zA-Z]{2,})$`
	data.ToolSpecificData.(map[string]interface{})["Explanation"] = "This regex validates an email address."

	view.Render(w, r, "regex-deconstructor.html", data)
}